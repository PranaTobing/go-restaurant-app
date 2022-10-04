package user

import "github.com/rocksus/go-restaurant-app/internal/model"

type Repository interface {
	RegisterUser(userData model.User) (model.User, error)
	CheckRegistered(username string) (bool, error)
	GenerateUserHash(password string) (hash string, err error)
	VerifyLogin(username, password string, userData model.User) (bool, error)
	GetUserData(username string) (model.User, error)
	CreateUserSession(userID string) (model.UserSession, error)
}
