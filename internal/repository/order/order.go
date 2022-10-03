package order

import (
	"github.com/rocksus/go-restaurant-app/internal/model"
	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &orderRepo{
		db: db,
	}
}

func (or *orderRepo) CreateOrder(order model.Order) (model.Order, error) {
	if err := or.db.Create(&order).Error; err != nil {
		return order, err
	}

	return order, nil
}

func (or *orderRepo) GetOrderInfo(orderID string) (model.Order, error) {
	var data model.Order

	if err := or.db.Where(model.Order{ID: orderID}).Preload("ProductOrders").First(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
