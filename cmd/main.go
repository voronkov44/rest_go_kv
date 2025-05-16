package main

import (
	"fmt"
	"net/http"
	"rest_go_kv/configs"
	"rest_go_kv/internal/auth"
	"rest_go_kv/internal/orders"
	"rest_go_kv/internal/users"
	"rest_go_kv/pkg/db"
	"rest_go_kv/pkg/jwt"
	"rest_go_kv/pkg/logger"

	"github.com/swaggo/http-swagger"
	_ "rest_go_kv/docs"
)

// @title           REST API
// @version         1.0
// @description     This is a server for managing users, authentication and orders.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Andrew Voronkov
// @contact.email  voronkovworkemail@gmail.com

// @host      localhost:8080
// @BasePath  /
func main() {
	// подключение логгера
	logger.InitLogger()

	//Чтение строки DSN
	conf := configs.LoadConfig()
	// Подключение к бд через горм
	database := db.NewDb(conf)
	router := http.NewServeMux()

	jwtManager := jwt.NewJWT(conf.Auth.Secret)

	// Подключение репозиториев
	userRepository := users.NewUserRepository(database)
	orderRepository := orders.NewOrderRepository(database)

	// Подключение всех хэндлеров
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:         conf,
		UserRepository: userRepository,
		JWT:            jwtManager,
	})
	users.NewUserHandler(router, users.UserHandlerDeps{
		UserRepository: userRepository,
	})
	orders.NewOrderHandler(router, orders.OrderHandlerDeps{
		OrderRepository: orderRepository,
		UserRepository:  userRepository,
	})

	router.Handle("/swagger/", httpSwagger.WrapHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server started at port 8080")
	logger.Info("Server started on port %d", 8080)
	server.ListenAndServe()
}
