package main

import (
	"fmt"
	"net/http"
	"rest_go_kv/configs"
	"rest_go_kv/internal/auth"
	"rest_go_kv/internal/orders"
	"rest_go_kv/internal/users"
	"rest_go_kv/pkg/db"
)

func main() {
	//Чтение строки DSN
	conf := configs.LoadConfig()
	// Подключение к бд через горм
	database := db.NewDb(conf)
	router := http.NewServeMux()

	// Подключение репозиториев
	userRepository := users.NewUserRepository(database)
	orderRepository := orders.NewOrderRepository(database)

	// Подключение всех хэндлеров
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:         conf,
		UserRepository: userRepository,
	})
	users.NewUserHandler(router, users.UserHandlerDeps{
		UserRepository: userRepository,
	})
	orders.NewOrderHandler(router, orders.OrderHandlerDeps{
		OrderRepository: orderRepository,
		UserRepository:  userRepository,
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server started at port 8080")
	server.ListenAndServe()
}
