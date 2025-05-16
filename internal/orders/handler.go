package orders

import (
	"net/http"
	"rest_go_kv/internal/users"
	"rest_go_kv/pkg/logger"
	"rest_go_kv/pkg/req"
	"rest_go_kv/pkg/res"
	"rest_go_kv/pkg/utils"
)

type OrderHandlerDeps struct {
	OrderRepository *OrderRepository
	UserRepository  *users.UserRepository
}

type OrderHandler struct {
	OrderRepository *OrderRepository
	UserRepository  *users.UserRepository
}

func NewOrderHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{
		OrderRepository: deps.OrderRepository,
		UserRepository:  deps.UserRepository,
	}
	router.HandleFunc("POST /users/{user_id}/orders", handler.Create())
	router.HandleFunc("GET /users/{user_id}/orders", handler.GoTo())
}

func (handler *OrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Received order creation request")

		userID, err := utils.ParseUserID(r)
		if err != nil {
			logger.Error("Invalid user ID: %v", err)
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		// Проверяем наличие юзера
		_, err = handler.UserRepository.GetById(userID)
		if err != nil {
			logger.Error("User with ID %d not found: %v", userID, err)
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		body, err := req.HandleBody[OrderCreateRequest](&w, r)
		if err != nil {
			logger.Error("Failed to parse order request body: %v", err)
			return
		}

		// Создание заказа
		order := NewOrder(userID, body.Product, uint(body.Quantity), body.Price)

		createdOrder, err := handler.OrderRepository.Create(order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logger.Error("Failed to create order for user %d: %v", userID, err)
			return
		}

		logger.Info("Order created: ID=%d, UserID=%d, Product=%s", createdOrder.ID, createdOrder.UserID, createdOrder.Product)
		res.Json(w, OrderResponse{
			ID:        createdOrder.ID,
			UserID:    createdOrder.UserID,
			Product:   createdOrder.Product,
			Quantity:  createdOrder.Quantity,
			Price:     createdOrder.Price,
			CreatedAt: createdOrder.CreatedAt,
		}, http.StatusOK)
	}
}

func (handler *OrderHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Received order list request")

		userID, err := utils.ParseUserID(r)
		if err != nil {
			logger.Error("Invalid user ID: %v", err)
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		// Проверяем наличие юзера
		_, err = handler.UserRepository.GetById(userID)
		if err != nil {
			logger.Error("User with ID %d not found: %v", userID, err)
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		// Получаем список заказов пользователя
		ordersList, err := handler.OrderRepository.GetByUserID(userID)
		if err != nil {
			logger.Error("Failed to get orders for user %d: %v", userID, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Info("Found %d orders for user ID %d", len(ordersList), userID)

		var response []OrderResponse
		for _, order := range ordersList {
			response = append(response, OrderResponse{
				ID:        order.ID,
				UserID:    order.UserID,
				Product:   order.Product,
				Quantity:  order.Quantity,
				Price:     order.Price,
				CreatedAt: order.CreatedAt,
			})
		}

		logger.Info("Found %d orders for user ID %d", len(ordersList), userID)
		res.Json(w, response, http.StatusOK)
	}
}
