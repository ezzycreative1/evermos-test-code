package middleware

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AuthMiddleware for Authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := c.Request.Header.Get("secret-key")
		env := os.Getenv("ENV")
		if env != "local" && govalidator.IsNull(secretKey) {
			if secretKey != os.Getenv("SECRET_KEY") {
				c.AbortWithStatus(401)
			}
		}
		c.Next()
	}
}

// TokenAuthMiddleware ..
// func TokenAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		err := auth.TokenValid(c)
// 		if err != nil {
// 			BaseHandler.RespondUnauthorized(c, "")
// 			return
// 		}
// 		c.Next()
// 	}
// }

// MaxAllowed specify max allowed connections
func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()
	}
}

// LoggerToFile ..
func LoggerToFile() gin.HandlerFunc {
	// log file
	fileName := path.Join("logging", "lintasarta.log")
	// write file
	//src, err := os.OpenFile("lintasarta.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	src, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err", err)
	}
	// instantiation
	logger := logrus.New()
	// Set output
	logger.Out = src
	// Set log level
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	// logger.AddHook(lfHook)
	return func(c *gin.Context) {
		// start time
		startTime := time.Now()
		// Processing request
		c.Next()
		// Stop time
		endTime := time.Now()
		// execution time
		latencyTime := endTime.Sub(startTime)
		// Request mode
		reqMethod := c.Request.Method
		// Request routing
		reqUri := c.Request.RequestURI
		// Status code
		statusCode := c.Writer.Status()
		// Request IP
		clientIP := c.ClientIP()
		// Log format
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
