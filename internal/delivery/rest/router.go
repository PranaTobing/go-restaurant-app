package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoadRoutes(e *echo.Echo, handler *handler) {
	menuGroup := e.Group("/menu")
	menuGroup.GET("", handler.GetMenu)

	orderGroup := e.Group("/order")
	orderGroup.POST("", handler.Order)
	orderGroup.GET("/:orderID", handler.GetOrderInfo)
}

func LoadMiddlewares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// the safety comes from using https, since Origin header is controlled by the browser
		AllowOrigins: []string{"https://restoku.com"},
	}))
}
