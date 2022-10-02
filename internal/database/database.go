package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB(dbAddress string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	seedDB(db)

	return db
}
