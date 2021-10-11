package category

import (
	BaseHandler "01-online-store/app/base/handler"
	"01-online-store/app/category"
	"01-online-store/models"
	"math"
	"strconv"

	"01-online-store/external/requests"
	resp "01-online-store/external/response"

	"github.com/gin-gonic/gin"
)

// CategoryHandler ..
type CategoryHandler struct {
	UseCase category.ICategoryUsecase
}

// GetAllCategoryHandler ..
func (ct *CategoryHandler) GetAllCategoryHandler(c *gin.Context) {
	queryPage := c.Query("page")
	queryLimit := c.Query("limit")
	limit := 10
	page := 1

	if len(queryPage) != 0 {
		i, _ := strconv.Atoi(queryPage)
		page = i
	}

	if len(queryLimit) != 0 {
		i, _ := strconv.Atoi(queryLimit)
		limit = i
	}

	count, listCat, err := ct.UseCase.GetCategory(c, page, limit)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	var pages resp.Paginate

	tpages := float64(count) / float64(limit)
	pages.TotalPage = math.Ceil(float64(tpages))
	pages.Count = int64(count)
	pages.Page = int64(page)
	type responsewithpage struct {
		Cat      *[]models.Category `json:"category"`
		Paginate resp.Paginate      `json:"pagination"`
	}

	res := responsewithpage{
		Cat:      listCat,
		Paginate: pages,
	}
	BaseHandler.RespondSuccess(c, "Success Get All Customer", res)
}

func (ct *CategoryHandler) GetCategoryByIDHandler(c *gin.Context) {
	custID := c.Param("id")
	convCustID, _ := strconv.ParseInt(custID, 10, 64)
	cust, err := ch.UseCase.GetCustomerByID(c, convCustID)
	if err != nil {
		if err != nil {
			BaseHandler.FailedResponseBackend(c, err)
			return
		}
	}
	BaseHandler.RespondSuccess(c, "success get customer", cust)
}

// CreateCustomerHandler ..
func (ch *CustomerHandler) CreateCustomerHandler(c *gin.Context) {
	requestBody := requests.CreateCustomerRequest{}
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	custID, err := ch.UseCase.CreateCustomer(c, requestBody)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	// Success
	BaseHandler.RespondSuccess(c, "create customer success", custID)
}

//EditCustomerHandler ..
func (ch *CustomerHandler) EditCustomerHandler(c *gin.Context) {
	requestBody := models.Customer{}
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	custID, err := ch.UseCase.EditCustomer(c, requestBody)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	// Success
	BaseHandler.RespondSuccess(c, "edit customer success", custID)
}
