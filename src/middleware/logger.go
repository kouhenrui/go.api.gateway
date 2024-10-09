package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.api.gateway/src/config"
	"time"
)

// 日志模块

// LoggerMiddleware logs incoming requests and their processing time
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Process request
		c.Next()

		// Log request details
		duration := time.Since(startTime)
		//log.Printf("Request: %s %s | Status: %d | Duration: %s",
		//	c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
		statusCode := c.Writer.Status()

		switch statusCode {
		case 200:
			config.Logger.WithFields(logrus.Fields{
				"status":   statusCode,
				"method":   c.Request.Method,
				"path":     c.Request.URL.Path,
				"duration": duration,
				"clientIP": c.ClientIP(),
			}).Info("Request completed successfully")
		case 404:
			config.Logger.WithFields(logrus.Fields{
				"status":   statusCode,
				"method":   c.Request.Method,
				"path":     c.Request.URL.Path,
				"duration": duration,
				"clientIP": c.ClientIP(),
			}).Warn("Resource not found")
		case 405:
			config.Logger.WithFields(logrus.Fields{
				"status":   statusCode,
				"method":   c.Request.Method,
				"path":     c.Request.URL.Path,
				"duration": duration,
				"clientIP": c.ClientIP(),
			}).Warn("Method not allowed")
		case 500:
			config.Logger.WithFields(logrus.Fields{
				"status":   statusCode,
				"method":   c.Request.Method,
				"path":     c.Request.URL.Path,
				"duration": duration,
				"clientIP": c.ClientIP(),
			}).Error("Internal server error")
		default:
			config.Logger.WithFields(logrus.Fields{
				"status":   statusCode,
				"method":   c.Request.Method,
				"path":     c.Request.URL.Path,
				"duration": duration,
				"clientIP": c.ClientIP(),
			}).Error("Request failed with an unexpected status code")
		}

	}
}
