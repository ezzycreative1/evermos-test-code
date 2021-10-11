package customer

import (
	"errors"
	"strconv"

	BaseHandler "01-online-store/app/base/handler"
	"01-online-store/app/customer"

	"01-online-store/external/requests"
	"01-online-store/helpers/response_mapping"

	"github.com/gin-gonic/gin"
)

// CustomerHandler ..
type CustomerHandler struct {
	UseCase customer.ICustomerUsecase
}

// CreateCustomerHandler ..
func (ch *CustomerHandler) CreateCustomerHandler(c *gin.Context) {
	requestBody := requests.CreateTicketRequest{}
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

// CountTicketHandler ..
func (th *TicketHandler) CountTicketHandler(c *gin.Context) {
	queryOrganization := c.Query("organization")
	queryUser := c.Query("user")
	queryProject := c.Query("project")
	queryTicketType := c.Query("ticket_type")
	queryStatus := c.Query("status")
	orgID := ""
	userID := ""
	projectID := ""
	tiketType := ""
	tStatus := ""

	if len(queryOrganization) != 0 {
		orgID = queryOrganization
	}

	if len(queryUser) != 0 {
		userID = queryUser
	}

	if len(queryProject) != 0 {
		projectID = queryProject
	}

	if len(queryTicketType) != 0 {
		tiketType = queryTicketType
	}

	if len(queryStatus) != 0 {
		tStatus = queryStatus
	}

	count, err := th.UseCase.GetCountTicket(c, userID, orgID, projectID, tiketType, tStatus)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	// Success
	BaseHandler.RespondSuccess(c, "count ticket success", count)
}

// GetTicketHandler ..
func (th *TicketHandler) GetTicketHandler(c *gin.Context) {
	queryOrganization := c.Query("organization")
	queryUser := c.Query("user")
	queryProject := c.Query("project")
	queryTicketType := c.Query("ticket_type")
	queryStatus := c.Query("status")
	queryPage := c.Query("page")
	queryLimit := c.Query("limit")
	orgID := ""
	userID := ""
	projectID := ""
	tiketType := ""
	tStatus := ""
	page := 0
	limit := 0

	if len(queryOrganization) != 0 {
		orgID = queryOrganization
	}

	if len(queryUser) != 0 {
		userID = queryUser
	}

	if len(queryProject) != 0 {
		projectID = queryProject
	}

	if len(queryTicketType) != 0 {
		tiketType = queryTicketType
	}

	if len(queryStatus) != 0 {
		tStatus = queryStatus
	}

	if len(queryPage) != 0 {
		i, _ := strconv.Atoi(queryPage)
		page = i
	}

	if len(queryLimit) != 0 {
		i, _ := strconv.Atoi(queryLimit)
		limit = i
	}

	tickets, err := th.UseCase.GetTickets(c, userID, orgID, projectID, tiketType, tStatus, page, limit)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	BaseHandler.RespondSuccess(c, "get all ticket success", tickets)
}

func (th *TicketHandler) GetTicketDetailHandler(c *gin.Context) {
	ticketID := c.Param("ticket_id")
	if ticketID == "" {
		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
			Error: errors.New("params ticket id required"),
			Code:  response_mapping.InvalidParam,
		})
		return
	}
	ticket, err := th.UseCase.GetTicketByID(c, ticketID)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	// Success
	BaseHandler.RespondSuccess(c, "get ticket "+ticketID+" success", ticket)
}

// UpdateStatusHandler ..
func (th *TicketHandler) UpdateStatusHandler(c *gin.Context) {
	requestBody := requests.UpdateStatusRequest{}
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	err = th.UseCase.UpdateStatus(c, requestBody)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	// Success
	BaseHandler.RespondSuccess(c, "Update ticket "+requestBody.TicketID+" success", nil)
}

// DeleteTicketHandler ..
func (th *TicketHandler) DeleteTicketHandler(c *gin.Context) {
	ticketID := c.Param("ticket_id")
	if ticketID == "" {
		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
			Error: errors.New("params ticket id required"),
			Code:  response_mapping.InvalidParam,
		})
		return
	}
	err := th.UseCase.DeleteTicket(c, ticketID)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	// Success
	BaseHandler.RespondSuccess(c, "Delete ticket "+ticketID+" success", nil)
}

// DeleteTicketPermanentHandler ..
func (th *TicketHandler) DeleteTicketPermanentHandler(c *gin.Context) {
	ticketID := c.Param("ticket_id")
	if ticketID == "" {
		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
			Error: errors.New("params ticket id required"),
			Code:  response_mapping.InvalidParam,
		})
		return
	}
	err := th.UseCase.DeleteTicketPermanent(c, ticketID)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	// Success
	BaseHandler.RespondSuccess(c, "Delete ticket "+ticketID+" success", nil)
}

// GetTicketTypeHandler ..
func (th *TicketHandler) GetTicketTypeHandler(c *gin.Context) {
	tType, err := th.UseCase.GetTicketType(c)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	BaseHandler.RespondSuccess(c, "success", tType)
	return
}

// GetRequestTypeHandler ..
func (th *TicketHandler) GetRequestTypeHandler(c *gin.Context) {
	rType, err := th.UseCase.GetRequestType(c)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	BaseHandler.RespondSuccess(c, "success", rType)
	return
}

// GetStatusHandler ..
func (th *TicketHandler) GetStatusHandler(c *gin.Context) {
	stTc, err := th.UseCase.GetStatus(c)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	BaseHandler.RespondSuccess(c, "success", stTc)
	return
}

func (th *TicketHandler) GetSiteHandler(c *gin.Context) {
	sites, err := th.UseCase.GetSites(c)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	BaseHandler.RespondSuccess(c, "success", sites)
	return
}

func (th *TicketHandler) GetImpactHandler(c *gin.Context) {
	impact, err := th.UseCase.GetImpact(c)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	BaseHandler.RespondSuccess(c, "success", impact)
	return
}

// GetUrgencyHandler ..
func (th *TicketHandler) GetUrgencyHandler(c *gin.Context) {
	urgent, err := th.UseCase.GetUrgency(c)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	BaseHandler.RespondSuccess(c, "success", urgent)
	return
}

// GetPriorityHandler ..
func (th *TicketHandler) GetPriorityHandler(c *gin.Context) {
	priority, err := th.UseCase.GetPriority(c)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	BaseHandler.RespondSuccess(c, "success", priority)
	return
}

// GetCategoryHandler ..
func (th *TicketHandler) GetCategoryHandler(c *gin.Context) {
	category, err := th.UseCase.GetCategory(c)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	BaseHandler.RespondSuccess(c, "success", category)
	return
}

// GetSubcategoryHandler ..
func (th *TicketHandler) GetSubcategoryHandler(c *gin.Context) {
	subcategory, err := th.UseCase.GetSubcategory(c)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	BaseHandler.RespondSuccess(c, "success", subcategory)
	return
}

// GetItemHandler ..
func (th *TicketHandler) GetItemHandler(c *gin.Context) {
	item, err := th.UseCase.GetItem(c)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	BaseHandler.RespondSuccess(c, "success", item)
	return
}
