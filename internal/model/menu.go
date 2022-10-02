package model

import (
	"github.com/rocksus/go-restaurant-app/internal/model/constant"
)

type MenuItem struct {
	OrderCode string `gorm:"primaryKey"`
	Name      string
	Price     int64
	Type      constant.MenuType
}
