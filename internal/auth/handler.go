package auth

import (
	"net/http"
	"rest_go_kv/configs"
	"rest_go_kv/internal/users"
	"rest_go_kv/pkg/jwt"
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

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}

		// Получаем юзера по email
		user, err := handler.UserRepository.GetByEmail(body.Email)
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		// Проверяем пароль
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
		if err != nil {
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

		res.Json(w, LoginResponse{
			Token: token,
		}, http.StatusOK)
	}
}
