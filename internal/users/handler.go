package users

import (
	"net/http"
	"rest_go_kv/pkg/logger"
	"rest_go_kv/pkg/req"
	"rest_go_kv/pkg/res"
	"rest_go_kv/pkg/utils"
	"strconv"
)

type UserHandlerDeps struct {
	UserRepository *UserRepository
}

type UserHandler struct {
	UserRepository *UserRepository
}

func NewUserHandler(router *http.ServeMux, deps UserHandlerDeps) {
	handler := &UserHandler{
		UserRepository: deps.UserRepository,
	}
	router.HandleFunc("POST /users", handler.Create())
	router.HandleFunc("GET /users", handler.GetAll())
	router.HandleFunc("GET /users/{id}", handler.GoTo())
	router.HandleFunc("PUT /users/{id}", handler.UpdateAll())
	router.HandleFunc("DELETE /users/{id}", handler.Delete())
}

// Create godoc
// @Summary Создать пользователя
// @Description Создаёт нового пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param request body UserCreateRequest true "Данные пользователя"
// @Success 200 {object} users.UserCreateResponse
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal error"
// @Router /users [post]
func (handler *UserHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Received request to create user")

		// Парсим тело запроса
		body, err := req.HandleBody[UserCreateRequest](&w, r)
		if err != nil {
			logger.Error("Failed to parse request body: %v", err)
			return
		}
		logger.Debug("Parsed request body: %+v", body)

		// Проверка на email
		exists, err := handler.UserRepository.IsEmailExist(body.Email)
		if err != nil {
			logger.Error("Failed to check email existence: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if exists {
			logger.Info("Attempt to create user with existing email: %s", body.Email)
			http.Error(w, "User already exists", http.StatusBadRequest)
			return
		}

		// Хешируем пароль
		hashedPassword, err := utils.HashPassword(body.Password)
		if err != nil {
			logger.Error("Failed to hash password: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Создаем пользователя
		user := NewUser(body.Name, body.Email, body.Age)
		user.Password = hashedPassword

		createdUser, err := handler.UserRepository.Create(user)
		if err != nil {
			logger.Error("Failed to create user: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Info("User created successfully: ID=%d, Email=%s", createdUser.ID, createdUser.Email)
		res.Json(w, UserCreateResponse{
			ID:    createdUser.ID,
			Name:  createdUser.Name,
			Email: createdUser.Email,
			Age:   createdUser.Age,
		}, http.StatusOK)
	}
}

// GetAll godoc
// @Summary Получить всех пользователей
// @Description Возвращает список всех пользователей
// @Tags users
// @Produce json
// @Success 200 {array} users.UserCreateResponse
// @Failure 500 {string} string "internal error"
// @Router /users [get]
func (handler *UserHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Received request to get all users")

		query := r.URL.Query()

		page, err := strconv.Atoi(query.Get("page"))
		if err != nil || page <= 0 {
			page = 1
		}

		limit, err := strconv.Atoi(query.Get("limit"))
		if err != nil || limit <= 0 {
			limit = 10
		}

		minAge, _ := strconv.Atoi(query.Get("min_age"))
		maxAge, _ := strconv.Atoi(query.Get("max_age"))

		logger.Debug("Query params: page=%d, limit=%d, minAge=%d, maxAge=%d", page, limit, minAge, maxAge)

		// Получаем пользователей из репозитория
		users, err := handler.UserRepository.GetAll(page, limit, minAge, maxAge)
		if err != nil {
			logger.Error("Failed to get users: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Info("Successfully retrieved %d users", len(users))

		// Маппим в респонс структуру (чтобы не вывести пользователю ненужные поля)
		var response []UserCreateResponse
		for _, u := range users {
			response = append(response, UserCreateResponse{
				ID:    u.ID,
				Name:  u.Name,
				Email: u.Email,
				Age:   u.Age,
			})
		}
		res.Json(w, response, http.StatusOK)
	}
}

// GoTo godoc
// @Summary Получить пользователя по ID
// @Description Возвращает пользователя по его ID
// @Tags users
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} users.UserCreateResponse
// @Failure 404 {string} string "user not found"
// @Router /users/{id} [get]
func (handler *UserHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Received request to get user by ID")

		// Получение id из path value
		id, err := utils.ParseID(r)
		if err != nil {
			logger.Error("Invalid user ID: %v", err)
			http.Error(w, "invalid user id", http.StatusBadRequest)
			return
		}
		logger.Debug("Parsed user ID: %d", id)

		// Получаем пользователя из репозитория
		user, err := handler.UserRepository.GetById(id)
		if err != nil {
			logger.Error("User not found: %v", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		logger.Info("User retrieved: ID=%d, Email=%s", user.ID, user.Email)

		// Маппинг в респонс структуру
		response := UserCreateResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Age:   user.Age,
		}
		res.Json(w, response, http.StatusOK)

	}
}

// UpdateAll godoc
// @Summary Обновить данные пользователя
// @Description Полностью обновляет данные пользователя по ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param request body UserUpdateRequest true "Обновлённые данные пользователя"
// @Success 200 {object} users.UserUpdateResponse
// @Failure 404 {string} string "user not found"
// @Failure 400 {string} string "bad request"
// @Router /users/{id} [put]
func (handler *UserHandler) UpdateAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Received request to update user")

		// Получение id из path value
		id, err := utils.ParseID(r)
		if err != nil {
			logger.Error("Invalid user ID: %v", err)
			http.Error(w, "invalid user id", http.StatusBadRequest)
			return
		}
		logger.Debug("Parsed user ID: %d", id)

		// Проверяем наличие пользователя
		existingUser, err := handler.UserRepository.GetById(id)
		if err != nil {
			logger.Error("User not found: %v", err)
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		// Парсим тело запроса
		body, err := req.HandleBody[UserUpdateRequest](&w, r)
		if err != nil {
			logger.Error("Failed to parse request body: %v", err)
			return
		}
		logger.Debug("Parsed request body: %+v", body)

		// Проверка на email (если поменялся)
		if existingUser.Email != body.Email {
			exists, err := handler.UserRepository.IsEmailExist(body.Email)
			if err != nil {
				logger.Error("Failed to check email existence: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if exists {
				logger.Info("Attempt to update user with existing email: %s", body.Email)
				http.Error(w, "User already exists", http.StatusBadRequest)
				return
			}
		}

		// Обновляем все поля
		existingUser.Name = body.Name
		existingUser.Email = body.Email
		existingUser.Age = body.Age

		// Хешируем пароль (в тз видел что этого нет, но решил добавить обновление и пароля тоже (необязательное поле (тз не противоречит), использую другую структурку))
		if body.Password != "" {
			hashedPassword, err := utils.HashPassword(body.Password)
			if err != nil {
				logger.Error("Failed to hash password: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			existingUser.Password = hashedPassword
		}

		// Сохраняем изменения
		updatedUser, err := handler.UserRepository.Update(existingUser)
		if err != nil {
			logger.Error("Failed to update user: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Info("User updated successfully: ID=%d", updatedUser.ID)

		// Возвращаем ответ
		res.Json(w, UserUpdateResponse{
			ID:    updatedUser.ID,
			Name:  updatedUser.Name,
			Email: updatedUser.Email,
			Age:   updatedUser.Age,
		}, http.StatusOK)
	}
}

// Delete godoc
// @Summary Удалить пользователя
// @Description Удаляет пользователя по ID
// @Tags users
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 204 {string} string "No Content"
// @Failure 404 {string} string "user not found"
// @Router /users/{id} [delete]
func (handler *UserHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Received request to delete user")

		// Получение id из path value
		id, err := utils.ParseID(r)
		if err != nil {
			logger.Error("Invalid user ID: %v", err)
			http.Error(w, "invalid user id", http.StatusBadRequest)
			return
		}
		logger.Debug("Parsed user ID: %d", id)

		_, err = handler.UserRepository.GetById(id)
		if err != nil {
			logger.Error("User not found: %v", err)
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		// Удаляем пользователя
		err = handler.UserRepository.Delete(int(id))
		if err != nil {
			logger.Error("Failed to delete user: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Info("User deleted successfully: ID=%d", id)

		w.WriteHeader(http.StatusNoContent)
	}
}
