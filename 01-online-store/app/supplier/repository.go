package supplier

import (
	"01-online-store/models"
)

type Repository interface {
	CountSupp() (int64, error)
	GetAllSupplier(page int, limit int) (supp []models.Supplier, err error)
	GetSupplierByID(suppID int64) (supp models.Supplier, err error)
	CreateSupplier(supp *models.Supplier) (*int64, error)
	UpdateSupplier(supp *models.Supplier) (*int64, error)
}
