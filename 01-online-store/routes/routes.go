package app

import (
	HCHandler "01-online-store/app/healthcheck/handler"

	"github.com/gin-gonic/gin"
)

// HealthCheckHTTPHandler routes
func HealthCheckHTTPHandler(router *gin.Engine) {
	handler := &HCHandler.HealthCheckHandler{}
	router.GET("/check", handler.Check)
}
