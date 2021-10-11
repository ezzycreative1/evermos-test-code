package repository

import (
	"01-online-store/app/supplier"
	"01-online-store/models"

	"github.com/jinzhu/gorm"
)

type supplierRepository struct {
	Conn *gorm.DB
}

func NewSupplierRepository(Conn *gorm.DB) supplier.Repository {
	return &supplierRepository{Conn}
}

func (sr *supplierRepository) CountSupp() (int64, error) {
	query := sr.Conn

	var countdata int64
	err := query.Model(&models.Supplier{}).Count(&countdata).Error
	return countdata, err
}

func (sr *supplierRepository) GetAllSupplier(page int, limit int) (supp []models.Supplier, err error) {
	offset := (page - 1) * limit
	query := sr.Conn.Model(&models.Supplier{})

	if err = query.Where("deleted_at is NULL").Order("created_at desc").Limit(limit).Offset(offset).Find(&supp).Error; err != nil {
		return nil, err
	}
	return supp, nil
}

func (sr *supplierRepository) GetSupplier() (supp []models.Supplier, err error) {
	return supp, sr.Conn.First(&supp).Error
}

func (sr *supplierRepository) GetSupplierByID(custID int64) (supp models.Supplier, err error) {
	return supp, sr.Conn.Where("id = ?", custID).First(&supp).Error
}

func (sr *supplierRepository) CreateSupplier(supp *models.Supplier) (*int64, error) {
	if err := sr.Conn.Create(&supp).Error; err != nil {
		return nil, err
	}

	return &supp.ID, nil
}

func (sr *supplierRepository) UpdateSupplier(supp *models.Supplier) (*int64, error) {
	if err := sr.Conn.Update(&supp).Error; err != nil {
		return nil, err
	}

	return &supp.ID, nil
}
