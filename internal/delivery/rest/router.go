package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoadRoutes(e *echo.Echo, handler *handler) {
	menuGroup := e.Group("/menu", middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://restoku.com"},
	}))
	menuGroup.GET("", handler.GetMenu)

	orderGroup := e.Group("/order")
	orderGroup.POST("", handler.Order)
	orderGroup.GET("/:orderID", handler.GetOrderInfo)

	userGroup := e.Group("/user")
	userGroup.POST("/register", handler.RegisterUser)
}

func LoadMiddlewares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// the safety comes from using https, since Origin header is controlled by the browser
		AllowOrigins: []string{"https://restoku.com"},
	}))
}
