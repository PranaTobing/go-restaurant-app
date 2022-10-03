package rest

import (
	"github.com/labstack/echo/v4"
)

func LoadRoutes(e *echo.Echo, handler *handler) {
	menuGroup := e.Group("/menu")
	menuGroup.GET("", handler.GetMenu)

	orderGroup := e.Group("/order")
	orderGroup.POST("", handler.Order)
	orderGroup.GET("/:orderID", handler.GetOrderInfo)
}
