package order

import (
	"context"

	"github.com/rocksus/go-restaurant-app/internal/model"
)

type Repository interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GetOrderInfo(ctx context.Context, orderID string) (model.Order, error)
}
