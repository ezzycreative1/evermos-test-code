package supplier

import (
	"01-online-store/external/requests"
	"01-online-store/external/response"
	"01-online-store/models"

	"github.com/gin-gonic/gin"
)

type ISupplierUsecase interface {
	GetSupplier(c *gin.Context, page int, limit int) (int64, *[]models.Supplier, error)
	GetSupplierByID(c *gin.Context, customerID int64) (*models.Supplier, error)
	CreateSupplier(c *gin.Context, req requests.CreateSupplierRequest) (*response.SupplierResp, error)
	EditSupplier(c *gin.Context, req models.Supplier) (*response.SupplierResp, error)
}
