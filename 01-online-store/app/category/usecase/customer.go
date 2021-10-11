package usecase

import (
	"01-online-store/app/category"
	"01-online-store/external/requests"
	"01-online-store/external/response"
	"01-online-store/models"

	"github.com/gin-gonic/gin"
)

type categoryUsecase struct {
	repo category.Repository
}

func NewCategoryUsecase(repo category.Repository) category.ICategoryUsecase {
	return &categoryUsecase{
		repo: repo,
	}
}

func (cu *categoryUsecase) GetCategory(c *gin.Context, page int, limit int) (int64, *[]models.Category, error) {
	count, err := cu.repo.CountCat()
	if err != nil {
		return count, nil, err
	}
	listCat, err := cu.repo.GetAllCategory(page, limit)
	if err != nil {
		return count, nil, err
	}
	return count, &listCat, nil
}

func (cu *categoryUsecase) GetCategoryByID(c *gin.Context, categoryID int64) (*models.Category, error) {
	cat, err := cu.repo.GetCategoryByID(categoryID)
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

func (cu *categoryUsecase) CreateCategory(c *gin.Context, req requests.CreateCategoryRequest) (*response.CategoryResp, error) {
	catID, err := cu.repo.CreateCategory(
		&models.Category{
			Name: req.Name,
		},
	)
	if err != nil {
		return nil, err
	}

	response := &response.CategoryResp{}
	response.CategoryCode = *catID

	return response, nil
}

func (cu *categoryUsecase) EditCategory(c *gin.Context, req models.Category) (*response.CategoryResp, error) {
	catID, err := cu.repo.UpdateCategory(
		&models.Category{
			Name: req.Name,
		},
	)
	if err != nil {
		return nil, err
	}

	response := &response.CategoryResp{}
	response.CategoryCode = *catID

	return response, nil
}
