package main

import (
	"github.com/labstack/echo/v4"

	mRepo "github.com/rocksus/go-restaurant-app/internal/repository/menu"
	oRepo "github.com/rocksus/go-restaurant-app/internal/repository/order"
	uRepo "github.com/rocksus/go-restaurant-app/internal/repository/user"
	rUsecase "github.com/rocksus/go-restaurant-app/internal/usecase/resto"

	"github.com/rocksus/go-restaurant-app/internal/database"
	"github.com/rocksus/go-restaurant-app/internal/delivery/rest"
)

func main() {
	e := echo.New()

	db := database.GetDB("host=localhost port=5432 user=postgres password=postgres dbname=go_resto sslmode=disable")
	secret := "AES256Key-32Characters1234567890"

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32)
	if err != nil {
		panic(err)
	}
	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo, userRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadMiddlewares(e)
	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start((":14045")))
}
