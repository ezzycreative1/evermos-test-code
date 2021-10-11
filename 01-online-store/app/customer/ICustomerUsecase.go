package customer

import (
	"01-online-store/external/requests"
	"01-online-store/external/response"
	"01-online-store/models"

	"github.com/gin-gonic/gin"
)

type ICustomerUsecase interface {
	GetCustomer(c *gin.Context, page int, limit int) (int64, *[]models.Customer, error)
	GetCustomerByID(c *gin.Context, customerID int64) (*models.Customer, error)
	CreateCustomer(c *gin.Context, req requests.CreateCustomerRequest) (*response.CustomerResp, error)
	EditCustomer(c *gin.Context, req models.Customer) (*response.CustomerResp, error)
}
