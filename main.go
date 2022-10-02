package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/menu/food", get_food_menu)
	e.GET("/menu/drinks", get_drinks_menu)

	e.Logger.Fatal(e.Start((":14045")))
}

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int64
}

func get_food_menu(c echo.Context) error {
	food_menu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": food_menu,
	})
}

func get_drinks_menu(c echo.Context) error {
	drinks_menu := []MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     4000,
		},
		{
			Name:      "Es Teh Manis",
			OrderCode: "es_teh_manis",
			Price:     5000,
		},
		{
			Name:      "Air Mineral",
			OrderCode: "air_mineral",
			Price:     7000,
		},
		{
			Name:      "Jus Apel",
			OrderCode: "jus_apel",
			Price:     14000,
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": drinks_menu,
	})
}
