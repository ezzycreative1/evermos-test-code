package customer

import (
	"01-online-store/models"
)

type Repository interface {
	CountCust() (int64, error)
	GetAllCustomer(page int, limit int) (cust []models.Customer, err error)
	GetCustomerByID(custID int64) (cust models.Customer, err error)
	CreateCustomer(cust *models.Customer) (*int64, error)
	UpdateCustomer(cust *models.Customer) (*int64, error)
}
