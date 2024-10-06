package middleware

import (
	"github.com/gin-gonic/gin"
	"go.api.gateway/src/config"
	"net/http"
)

var e = config.CasbinEnforcer{}

// CasbinMiddleware 中间件
func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户、方法和路径
		user := c.GetString("user") // 假设已经通过其他中间件设置了用户名
		obj := c.Request.URL.Path   // 访问路径
		act := c.Request.Method     // 请求方法

		// 检查用户权限
		ok, err := e.CheckPermission(user, obj, act)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "内部错误"})
			return
		}
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "无权访问"})
			return
		}
		c.Next()
	}
}
