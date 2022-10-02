package order

import (
	"github.com/google/uuid"
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

func (or *orderRepo) CreateOrder(order model.OrderData) (string, error) {
	if order.ID == "" {
		order.ID = uuid.New().String()
	}
	or.db.Create(&order)
	return order.ID, nil
}
