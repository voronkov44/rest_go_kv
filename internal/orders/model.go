package orders

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID   uint    `json:"user_id" gorm:"user_id"`
	Product  string  `json:"product" gorm:"product"`
	Quantity uint    `json:"quantity" gorm:"quantity"`
	Price    float64 `json:"price" gorm:"price"`
}

func NewOrder(user_id uint, product string, quantity uint, price float64) *Order {
	return &Order{
		UserID:   user_id,
		Product:  product,
		Quantity: quantity,
		Price:    price,
	}
}
