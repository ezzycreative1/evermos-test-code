package repository

import (
	"01-online-store/app/category"
	"01-online-store/models"

	"github.com/jinzhu/gorm"
)

type categoryRepository struct {
	Conn *gorm.DB
}

func NewCategoryRepository(Conn *gorm.DB) category.Repository {
	return &categoryRepository{Conn}
}

func (cr *categoryRepository) CountCat() (int64, error) {
	query := cr.Conn

	var countdata int64
	err := query.Model(&models.Category{}).Count(&countdata).Error
	return countdata, err
}

func (cr *categoryRepository) GetAllCategory(page int, limit int) (cat []models.Category, err error) {
	offset := (page - 1) * limit
	query := cr.Conn.Model(&models.Category{})

	if err = query.Where("deleted_at is NULL").Order("created_at desc").Limit(limit).Offset(offset).Find(&cat).Error; err != nil {
		return nil, err
	}
	return cat, nil
}

func (cr *categoryRepository) GetCategory() (cat []models.Category, err error) {
	return cat, cr.Conn.First(&cat).Error
}

func (cr *categoryRepository) GetCategoryByID(catID int64) (cat models.Category, err error) {
	return cat, cr.Conn.Where("id = ?", catID).First(&cat).Error
}

func (cr *categoryRepository) CreateCategory(cat *models.Category) (*int64, error) {
	if err := cr.Conn.Create(&cat).Error; err != nil {
		return nil, err
	}

	return &cat.ID, nil
}

func (cr *categoryRepository) UpdateCategory(cat *models.Category) (*int64, error) {
	if err := cr.Conn.Update(&cat).Error; err != nil {
		return nil, err
	}

	return &cat.ID, nil
}
