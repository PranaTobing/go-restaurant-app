package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MenuType string

const (
	Food  MenuType = "food"
	Drink MenuType = "drink"
)

const (
	db_address = "host=localhost port=5432 user=postgres password=postgres dbname=go_resto sslmode=disable"
)

type MenuItem struct {
	OrderCode string `gorm:"primaryKey"`
	Name      string
	Price     int64
	Type      MenuType
}

func seed_db() {
	db_address := "host=localhost port=5432 user=postgres password=postgres dbname=go_resto sslmode=disable"
	db, err := gorm.Open(postgres.Open(db_address), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&MenuItem{})

	food_menu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      Food,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      Food,
		},
	}

	drinks_menu := []MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     4000,
			Type:      Drink,
		},
		{
			Name:      "Es Teh Manis",
			OrderCode: "es_teh_manis",
			Price:     5000,
			Type:      Drink,
		},
		{
			Name:      "Air Mineral",
			OrderCode: "air_mineral",
			Price:     7000,
			Type:      Drink,
		},
		{
			Name:      "Jus Apel",
			OrderCode: "jus_apel",
			Price:     14000,
			Type:      Drink,
		},
	}

	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		fmt.Println("Seeding db data...")
		db.Create(&food_menu)
		db.Create(&drinks_menu)
	}
}

func main() {
	e := echo.New()

	seed_db()

	e.GET("/menu", get_menu)

	e.Logger.Fatal(e.Start((":14045")))
}

func get_menu(c echo.Context) error {
	menu_type := c.FormValue("menu_type")

	db, err := gorm.Open(postgres.Open(db_address), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	menu_data := make([]MenuItem, 0)

	db.Where(MenuItem{Type: MenuType(menu_type)}).Find(&menu_data)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menu_data,
	})
}
