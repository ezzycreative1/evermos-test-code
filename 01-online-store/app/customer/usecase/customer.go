package usecase

import (
	"01-online-store/external/requests"
	"01-online-store/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ticketUsecase struct {
	repo ticket.Repository
}

func NewTicketUsecase(repo ticket.Repository) ticket.ITicketUsecase {
	return &ticketUsecase{
		repo: repo,
	}
}

func (tu *ticketUsecase) CreateTicket(c *gin.Context, req requests.CreateTicketRequest) (*models.TicketResponse, error) {
	u, err := uuid.NewRandom()
	tID := u.String()
	if err != nil {
		log.Println(fmt.Errorf("failed to generate UUID: %w", err))
		return nil, err
	}

	// param := requests.CreateTicketParam{}
	// param.Request.Requester.Name = req.Requester
	// param.Request.RequestType.Name = req.RequestType
	// param.Request.Status.Name = req.Status
	// param.Request.Mode.Name = "Service Portal"
	// param.Request.Technician.Name = "ITS Service Desk"
	// param.Request.Impact.Name = req.Impact
	// param.Request.Urgency.Name = req.Urgency
	// param.Request.Priority.Name = req.Priority
	// param.Request.Category.Name = req.Category
	// param.Request.Subcategory.Name = req.Subcategory
	// param.Request.Item.Name = req.Item
	// param.Request.Template.Name = "Incident Cloudeka Flexi"
	// param.Request.Subject = req.Subject
	// param.Request.Description = req.Description
	// param.Request.Site.Name = req.Site
	// ticket, err := tu.repo.SDPCreateTicket(param)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("ticket => ", ticket)
	// dataTicketType, err := tu.repo.GetTicketTypeByName(req.TicketType)
	// if err != nil {
	// 	return nil, err
	// }

	// dataRequestType, err := tu.repo.GetRequestTypeByName(req.RequestType)
	// if err != nil {
	// 	return nil, err
	// }

	// dataStatus, err := tu.repo.GetStatusByName(req.Status)
	// if err != nil {
	// 	return nil, err
	// }

	// dataImpact, err := tu.repo.GetImpactByName(req.Impact)
	// if err != nil {
	// 	return nil, err
	// }

	// dataUrgency, err := tu.repo.GetUrgencyByName(req.Urgency)
	// if err != nil {
	// 	return nil, err
	// }

	// dataPriority, err := tu.repo.GetPriorityByName(req.Priority)
	// if err != nil {
	// 	return nil, err
	// }

	// dataCategory, err := tu.repo.GetCategoryByName(req.Category)
	// if err != nil {
	// 	return nil, err
	// }

	// dataSubcategory, err := tu.repo.GetSubcategoryByName(req.Subcategory)
	// if err != nil {
	// 	return nil, err
	// }

	// dataItem, err := tu.repo.GetItemByName(req.Item)
	// if err != nil {
	// 	return nil, err
	// }

	// dataSite, err := tu.repo.GetSiteByName(req.Site)
	// if err != nil {
	// 	return nil, err
	// }

	//sdpID, _ := strconv.ParseInt(ticket.Request.ID, 10, 64)
	//sdpTechnicianID, _ := strconv.ParseInt(ticket.Request.Technician.ID, 10, 64)
	tiketID, err := tu.repo.InsertTicket(
		&models.Ticket{
			ID:              tID,
			OrganizationID:  req.OrganizationID,
			ProjectID:       req.ProjectID,
			UserID:          req.UserID,
			SDPTicketID:     0,
			SDPTechnicianID: 0,
			Subject:         req.Subject,
			Description:     req.Description,
			TicketTypeID:    req.TicketTypeID,
			RequestTypeID:   req.RequestTypeID,
			StatusID:        req.StatusID,
			ImpactID:        req.ImpactID,
			UrgencyID:       req.UrgencyID,
			PriorityID:      req.PriorityID,
			CategoryID:      req.CategoryID,
			SubcategoryID:   req.SubcategoryID,
			ItemID:          req.ItemID,
			SiteID:          req.SiteID,
		},
	)
	if err != nil {
		return nil, err
	}

	response := &models.TicketResponse{}
	response.TicketCode = *tiketID

	return response, nil
}
