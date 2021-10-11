package usecase

import (
	"01-online-store/app/supplier"
	"01-online-store/external/requests"
	"01-online-store/external/response"
	"01-online-store/models"

	"github.com/gin-gonic/gin"
)

type supplierUsecase struct {
	repo supplier.Repository
}

func NewSupplierUsecase(repo supplier.Repository) supplier.ISupplierUsecase {
	return &supplierUsecase{
		repo: repo,
	}
}

func (su *supplierUsecase) GetSupplier(c *gin.Context, page int, limit int) (int64, *[]models.Supplier, error) {
	count, err := su.repo.CountSupp()
	if err != nil {
		return count, nil, err
	}
	listSupp, err := su.repo.GetAllSupplier(page, limit)
	if err != nil {
		return count, nil, err
	}
	return count, &listSupp, nil
}

func (su *supplierUsecase) GetSupplierByID(c *gin.Context, supplierID int64) (*models.Supplier, error) {
	cust, err := su.repo.GetSupplierByID(supplierID)
	if err != nil {
		return nil, err
	}
	return &cust, nil
}

func (su *supplierUsecase) CreateSupplier(c *gin.Context, req requests.CreateSupplierRequest) (*response.SupplierResp, error) {
	suppID, err := su.repo.CreateSupplier(
		&models.Supplier{
			Name:        req.Name,
			Email:       req.Email,
			PhoneNumber: req.PhoneNumber,
			Address:     req.Address,
			Kota:        req.Kota,
			Status:      "active",
		},
	)
	if err != nil {
		return nil, err
	}

	response := &response.SupplierResp{}
	response.SupplierCode = *suppID

	return response, nil
}

func (su *supplierUsecase) EditSupplier(c *gin.Context, req models.Supplier) (*response.SupplierResp, error) {
	suppID, err := su.repo.UpdateSupplier(
		&models.Supplier{
			Name:        req.Name,
			Email:       req.Email,
			PhoneNumber: req.PhoneNumber,
			Address:     req.Address,
			Kota:        req.Kota,
			Status:      req.Status,
		},
	)
	if err != nil {
		return nil, err
	}

	response := &response.SupplierResp{}
	response.SupplierCode = *suppID

	return response, nil
}
