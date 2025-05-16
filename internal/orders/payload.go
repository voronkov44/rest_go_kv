package orders

import "time"

type OrderCreateRequest struct {
	Product  string  `json:"product" validate:"required" example:"Laptop"`
	Quantity int     `json:"quantity" validate:"required" example:"1"`
	Price    float64 `json:"price" validate:"required" example:"1599.99"`
}

type OrderResponse struct {
	ID        uint      `json:"id" example:"1"`
	UserID    uint      `json:"user_id" example:"12"`
	Product   string    `json:"product" example:"Laptop"`
	Quantity  uint      `json:"quantity" example:"1"`
	Price     float64   `json:"price" example:"1599.99"`
	CreatedAt time.Time `json:"created_at" example:"2025-05-15 12:00:00"`
}
