package usecase

import (
	"01-online-store/app/customer"
	"01-online-store/external/requests"
	"01-online-store/external/response"
	"01-online-store/models"

	"github.com/gin-gonic/gin"
)

type customerUsecase struct {
	repo customer.Repository
}

func NewCustomerUsecase(repo customer.Repository) customer.ICustomerUsecase {
	return &customerUsecase{
		repo: repo,
	}
}

func (cu *customerUsecase) GetCustomer(c *gin.Context, page int, limit int) (int64, *[]models.Customer, error) {
	count, err := cu.repo.CountCust()
	if err != nil {
		return count, nil, err
	}
	listCust, err := cu.repo.GetAllCustomer(page, limit)
	if err != nil {
		return count, nil, err
	}
	return count, &listCust, nil
}

func (cu *customerUsecase) GetCustomerByID(c *gin.Context, customerID int64) (*models.Customer, error) {
	cust, err := cu.repo.GetCustomerByID(customerID)
	if err != nil {
		return nil, err
	}
	return &cust, nil
}

func (cu *customerUsecase) CreateCustomer(c *gin.Context, req requests.CreateCustomerRequest) (*response.CustomerResp, error) {
	custID, err := cu.repo.CreateCustomer(
		&models.Customer{
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

	response := &response.CustomerResp{}
	response.CustomerCode = *custID

	return response, nil
}

func (cu *customerUsecase) EditCustomer(c *gin.Context, req models.Customer) (*response.CustomerResp, error) {
	custID, err := cu.repo.UpdateCustomer(
		&models.Customer{
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

	response := &response.CustomerResp{}
	response.CustomerCode = *custID

	return response, nil
}
