package resto

import (
	"github.com/rocksus/go-restaurant-app/internal/model"
	"github.com/rocksus/go-restaurant-app/internal/repository/menu"
)

type restoUsecase struct {
	menuRepo menu.Repository
}

func GetUsecase(menuRepo menu.Repository) Usecase {
	return &restoUsecase{
		menuRepo: menuRepo,
	}
}

func (m *restoUsecase) GetMenu(menuType string) ([]model.MenuItem, error) {
	return m.menuRepo.GetMenu(menuType)
}
