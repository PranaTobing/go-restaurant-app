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
	dbAddress = "host=localhost port=5432 user=postgres password=postgres dbname=go_resto sslmode=disable"
)

type MenuItem struct {
	OrderCode string `gorm:"primaryKey"`
	Name      string
	Price     int64
	Type      MenuType
}

func seedDB() {
	dbAddress := "host=localhost port=5432 user=postgres password=postgres dbname=go_resto sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&MenuItem{})

	foodMenu := []MenuItem{
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

	drinksMenu := []MenuItem{
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
		db.Create(&foodMenu)
		db.Create(&drinksMenu)
	}
}

func main() {
	e := echo.New()

	seedDB()

	e.GET("/menu", GetMenu)

	e.Logger.Fatal(e.Start((":14045")))
}

func GetMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	menuData := make([]MenuItem, 0)

	if err := db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}
