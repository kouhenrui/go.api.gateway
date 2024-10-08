package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.api.gateway/src/config/response"
	"net/http"
)

func MethodNotAllowedHandler(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, response.Error(http.StatusMethodNotAllowed, "方法不允许"))
}
func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, response.Error(http.StatusNotFound, "资源不存在"))
}

// CORSMiddleware 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		fmt.Println("跨域请求处理中间件")
		c.Next()
	}
}
