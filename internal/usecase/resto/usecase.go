package resto

import (
	"github.com/rocksus/go-restaurant-app/internal/model"
)

type Usecase interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	Order(request model.OrderMenuRequest) (model.Order, error)
	GetOrderInfo(request model.GetOrderInfoRequest) (model.Order, error)
	RegisterUser(request model.RegisterRequest) (model.User, error)
	Login(request model.LoginRequest) (model.UserSession, error)
	CheckSession(sessionData model.UserSession) (userID string, err error)
}
