package auth

import (
	"net/http"
	"rest_go_kv/configs"
	"rest_go_kv/internal/users"
	"rest_go_kv/pkg/jwt"
	"rest_go_kv/pkg/logger"
	"rest_go_kv/pkg/req"
	"rest_go_kv/pkg/res"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandlerDeps struct {
	*configs.Config
	UserRepository *users.UserRepository
	JWT            *jwt.JWT
}

type AuthHandler struct {
	*configs.Config
	UserRepository *users.UserRepository
	JWT            *jwt.JWT
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:         deps.Config,
		UserRepository: deps.UserRepository,
		JWT:            deps.JWT,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
}

// Login godoc
// @Summary Авторизация пользователя
// @Description Авторизация пользователя по email и паролю
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Данные для авторизации"
// @Success 200 {object} LoginResponse
// @Failure 401 {string} string "invalid credentials"
// @Failure 500 {string} string "could not generate token"
// @Router /auth/login [post]
func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Login attempt received")
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			logger.Error("Failed to parse login request body: %v", err)
			return
		}
		logger.Debug("Login request parsed successfully: %+v", body)

		// Получаем юзера по email
		user, err := handler.UserRepository.GetByEmail(body.Email)
		if err != nil {
			logger.Error("User not found for email: %s. Error: %v", body.Email, err)
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		logger.Info("User %s authenticated successfully", body.Email)

		// Проверяем пароль
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
		if err != nil {
			logger.Error("Failed to generate JWT token for user: %s. Error: %v", body.Email, err)
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		// Генерируем JWT
		token, err := handler.JWT.Create(jwt.JWTData{
			Email: user.Email,
		})
		if err != nil {
			http.Error(w, "could not generate token", http.StatusInternalServerError)
			return
		}

		logger.Debug("JWT token generated: %s", token)

		res.Json(w, LoginResponse{
			Token: token,
		}, http.StatusOK)
		logger.Info("Login response sent successfully for user: %s", body.Email)
	}
}
