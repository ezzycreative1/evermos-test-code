package repository

import (
	"01-online-store/app/customer"
	"01-online-store/models"

	"github.com/jinzhu/gorm"
)

type customerRepository struct {
	Conn *gorm.DB
}

func NewCustomerRepository(Conn *gorm.DB) customer.Repository {
	return &customerRepository{Conn}
}

func (ur *customerRepository) CountCust() (int64, error) {
	query := ur.Conn

	var countdata int64
	err := query.Model(&models.Customer{}).Count(&countdata).Error
	return countdata, err
}

func (ur *customerRepository) GetAllCustomer(page int, limit int) (cust []models.Customer, err error) {
	offset := (page - 1) * limit
	query := ur.Conn.Model(&models.Customer{})

	if err = query.Where("deleted_at is NULL").Order("created_at desc").Limit(limit).Offset(offset).Find(&cust).Error; err != nil {
		return nil, err
	}
	return cust, nil
}

func (ur *customerRepository) GetCustomer() (cust []models.Customer, err error) {
	return cust, ur.Conn.First(&cust).Error
}

func (ur *customerRepository) GetCustomerByID(custID int64) (cust models.Customer, err error) {
	return cust, ur.Conn.Where("id = ?", custID).First(&cust).Error
}

func (ur *customerRepository) CreateCustomer(cust *models.Customer) (*int64, error) {
	if err := ur.Conn.Create(&cust).Error; err != nil {
		return nil, err
	}

	return &cust.ID, nil
}

func (ur *customerRepository) UpdateCustomer(cust *models.Customer) (*int64, error) {
	if err := ur.Conn.Update(&cust).Error; err != nil {
		return nil, err
	}

	return &cust.ID, nil
}
