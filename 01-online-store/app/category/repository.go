package category

import (
	"01-online-store/models"
)

type Repository interface {
	CountCat() (int64, error)
	GetAllCategory(page int, limit int) (cat []models.Category, err error)
	GetCategoryByID(catID int64) (cat models.Category, err error)
	CreateCategory(cat *models.Category) (*int64, error)
	UpdateCategory(cat *models.Category) (*int64, error)
}
