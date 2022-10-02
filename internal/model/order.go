package model

import "github.com/rocksus/go-restaurant-app/internal/model/constant"

type OrderData struct {
	ID          string `gorm:"primaryKey"`
	ProductCode string
	Quantity    int
	TotalPrice  int64
	Status      constant.OrderStatus
}
