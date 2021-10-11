package category

import (
	"01-online-store/external/requests"
	"01-online-store/external/response"
	"01-online-store/models"

	"github.com/gin-gonic/gin"
)

type ICategoryUsecase interface {
	GetCategory(c *gin.Context, page int, limit int) (int64, *[]models.Category, error)
	GetCategoryByID(c *gin.Context, categoryID int64) (*models.Category, error)
	CreateCategory(c *gin.Context, req requests.CreateCategoryRequest) (*response.CategoryResp, error)
	EditCategory(c *gin.Context, req models.Category) (*response.CategoryResp, error)
}
