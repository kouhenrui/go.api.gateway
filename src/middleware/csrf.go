package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-csrf"
	"go.api.gateway/src/config"
	"go.api.gateway/src/config/response"
	"net/http"
)

// CSRFMiddleware CSRF 中间件，适用于需要保护的路由
func CSRFMiddleware() gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret: config.ViperConfig.Service.CSRFKey,
		ErrorFunc: func(c *gin.Context) {
			c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, "CSRF token mismatch")) // gin.H{"error": "CSRF token mismatch"})
			c.Abort()
		},
	})
}

// GetCSRFToken 获取 CSRF Token 的函数
func GetCSRFToken(c *gin.Context) string {
	return csrf.GetToken(c)
}
