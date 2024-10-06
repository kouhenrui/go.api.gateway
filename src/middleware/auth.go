package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var validAPIKey = "your-api-key-secret"

// AuthJWTMiddleware  verifies the JWT token
func AuthJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		//tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		//token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		//	return []byte("your_jwt_secret"), nil
		//})

		//if err != nil || !token.Valid {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		//	c.Abort()
		//	return
		//}

		c.Next()
	}
}

// AuthCOOKIEMiddleware 验证 Cookie 是否存在并有效
func AuthCOOKIEMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session_id")
		if err != nil || cookie == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "未登录"})
			return
		}

		// 验证 session_id 的合法性
		// 假设session_id存储在Redis或数据库中
		// 在此处查询和验证 session_id

		c.Next()
	}
}

// APIKeyAuthMiddleware API 密钥验证中间件
func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey == "" || apiKey != validAPIKey {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "API key is missing or invalid",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
