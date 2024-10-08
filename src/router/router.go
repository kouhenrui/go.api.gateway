package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.api.gateway/src/config/response"
	"go.api.gateway/src/middleware"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.MaxMultipartMemory = 20 << 20
	r.Use(middleware.CORSMiddleware())        //处理跨域请求
	r.Use(middleware.LoggerMiddleware())      //处理日志记录
	r.Use(middleware.MethodNotAllowedHandler) //处理405请求
	r.Use(middleware.NotFoundHandler)         //处理404请求
	// Middleware for JWT authentication
	//r.Use(middleware.AuthJWTMiddleware()) //处理token请求头鉴定

	//r.Use(middleware.AuthCOOKIEMiddleware()) //处理cookie

	// Middleware for rate limiting
	//r.Use(middleware.RateLimiterMiddleware()) //处理限流请求

	//r.Use(middleware.CacheMiddleware()) //处理缓存请求

	r.Use(middleware.ErrorHandler()) //处理全局异常捕捉

	// Route for user service
	//r.Any("/user/*proxyPath", func(c *gin.Context) {
	//	middleware.ProxyRequest(c, "user_service")
	//})
	// 示例路由
	r.GET("/success", func(c *gin.Context) {
		data := map[string]string{
			"message": "请求成功",
		}
		c.JSON(http.StatusOK, response.Success(data))
	})

	r.GET("/error", func(c *gin.Context) {
		// 手动返回错误，测试全局捕获
		c.Error(
			gin.Error{
				Err:  fmt.Errorf("发生了某种错误"),
				Type: gin.ErrorTypePublic,
			})
		return
	})

	return r
}
