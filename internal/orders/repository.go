package orders

import "rest_go_kv/pkg/db"

type OrderRepository struct {
	Database *db.Db
}

func NewOrderRepository(database *db.Db) *OrderRepository {
	return &OrderRepository{
		Database: database,
	}
}

func (repo *OrderRepository) Create(Order *Order) (*Order, error) {
	result := repo.Database.DB.Create(Order)
	if result.Error != nil {
		return nil, result.Error
	}
	return Order, nil
}

func (repo *OrderRepository) GetByUserID(userID uint) ([]Order, error) {
	var orders []Order
	result := repo.Database.DB.Where("user_id = ?", userID).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
