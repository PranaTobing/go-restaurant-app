package order

import (
	"github.com/rocksus/go-restaurant-app/internal/model"
)

type Repository interface {
	CreateOrder(order model.OrderData) (string, error)
}
