package main

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/rocksus/go-restaurant-app/internal/logger"
	mRepo "github.com/rocksus/go-restaurant-app/internal/repository/menu"
	oRepo "github.com/rocksus/go-restaurant-app/internal/repository/order"
	uRepo "github.com/rocksus/go-restaurant-app/internal/repository/user"
	"github.com/rocksus/go-restaurant-app/internal/tracing"
	rUsecase "github.com/rocksus/go-restaurant-app/internal/usecase/resto"

	"github.com/rocksus/go-restaurant-app/internal/database"
	"github.com/rocksus/go-restaurant-app/internal/delivery/rest"
)

func main() {
	logger.Init()
	tracing.Init("http://localhost:14268/api/traces")

	e := echo.New()

	db := database.GetDB("host=localhost port=5432 user=postgres password=postgres dbname=go_resto sslmode=disable")
	secret := "AES256Key-32Characters1234567890"
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32, signKey, 60*time.Second)
	if err != nil {
		panic(err)
	}

	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo, userRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadMiddlewares(e)
	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start((":14045")))
}
