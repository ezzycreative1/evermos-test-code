package supplier

import (
	BaseHandler "01-online-store/app/base/handler"
	"01-online-store/app/supplier"
	"01-online-store/models"
	"math"
	"strconv"

	"01-online-store/external/requests"
	resp "01-online-store/external/response"

	"github.com/gin-gonic/gin"
)

// SupplierHandler ..
type SupplierHandler struct {
	UseCase supplier.ISupplierUsecase
}

// GetAllSupplierHandler ..
func (sh *SupplierHandler) GetAllSupplierHandler(c *gin.Context) {
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

	count, listSupp, err := sh.UseCase.GetSupplier(c, page, limit)
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
		Supp     *[]models.Supplier `json:"supplier"`
		Paginate resp.Paginate      `json:"pagination"`
	}

	res := responsewithpage{
		Supp:     listSupp,
		Paginate: pages,
	}
	BaseHandler.RespondSuccess(c, "Success Get All Supplier", res)
}

// GetSupplierByIDHandler ..
func (sh *SupplierHandler) GetSupplierByIDHandler(c *gin.Context) {
	suppID := c.Param("id")
	convSuppID, _ := strconv.ParseInt(suppID, 10, 64)
	cust, err := sh.UseCase.GetSupplierByID(c, convSuppID)
	if err != nil {
		if err != nil {
			BaseHandler.FailedResponseBackend(c, err)
			return
		}
	}
	BaseHandler.RespondSuccess(c, "success get supplier", cust)
}

// CreateSupplierHandler ..
func (sh *SupplierHandler) CreateSupplierHandler(c *gin.Context) {
	requestBody := requests.CreateSupplierRequest{}
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	custID, err := sh.UseCase.CreateSupplier(c, requestBody)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	// Success
	BaseHandler.RespondSuccess(c, "create supplier success", custID)
}

//EditCustomerHandler ..
func (sh *SupplierHandler) EditSupplierHandler(c *gin.Context) {
	requestBody := models.Supplier{}
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	custID, err := sh.UseCase.EditSupplier(c, requestBody)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	// Success
	BaseHandler.RespondSuccess(c, "edit supplier success", custID)
}
