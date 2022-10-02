package menu

import (
	"github.com/rocksus/go-restaurant-app/internal/model"
	"github.com/rocksus/go-restaurant-app/internal/model/constant"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &menuRepo{
		db: db,
	}
}

func (m *menuRepo) GetMenu(menuType string) ([]model.MenuItem, error) {

	menuData := make([]model.MenuItem, 0)

	m.db.Where(model.MenuItem{Type: constant.MenuType(menuType)}).Find(&menuData)

	return menuData, nil
}
