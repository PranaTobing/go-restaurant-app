package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *handler) GetMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	menuData, err := h.restoUsecase.GetMenuList(menuType)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][menu_handler][GetMenu] unable to get menu list")

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}
