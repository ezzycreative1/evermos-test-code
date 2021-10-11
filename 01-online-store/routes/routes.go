package app

import (
	"01-online-store/app/customer"
	CSHandler "01-online-store/app/customer/handler"
	HCHandler "01-online-store/app/healthcheck/handler"
	"01-online-store/app/supplier"
	SPHandler "01-online-store/app/supplier/handler"

	"github.com/gin-gonic/gin"
)

// HealthCheckHTTPHandler routes
func HealthCheckHTTPHandler(router *gin.Engine) {
	handler := &HCHandler.HealthCheckHandler{}
	router.GET("/check", handler.Check)
}

func CustomerHTTPHandler(router *gin.Engine, usecase customer.ICustomerUsecase) {
	handler := &CSHandler.CustomerHandler{UseCase: usecase}
	router.POST("/customer", handler.CreateCustomerHandler)
	router.PUT("/customer", handler.EditCustomerHandler)
	router.GET("/customer/:id", handler.GetCustomerByIDHandler)
	router.GET("/customer", handler.GetAllCustomerHandler)
}

func SupplierHTTPHandler(router *gin.Engine, usecase supplier.ISupplierUsecase) {
	handler := &SPHandler.SupplierHandler{UseCase: usecase}
	router.POST("/supplier", handler.CreateSupplierHandler)
	router.PUT("/supplier", handler.EditSupplierHandler)
	router.GET("/supplier/:id", handler.GetSupplierByIDHandler)
	router.GET("/supplier", handler.GetAllSupplierHandler)
}
