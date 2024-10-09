package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.api.gateway/pkg"
	"go.api.gateway/src/config/response"
	"go.api.gateway/src/middleware"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.MaxMultipartMemory = 20 << 20
	r.Use(middleware.CORSMiddleware())             //处理跨域请求
	r.Use(middleware.LoggerMiddleware())           //处理日志记录
	r.NoMethod(middleware.MethodNotAllowedHandler) //处理405请求
	r.NoRoute(middleware.NotFoundHandler)          //处理404请求
	r.Use(middleware.ErrorHandler())               //处理全局异常捕捉
	r.Use(middleware.AuthJWTMiddleware())          //处理token验证请求
	r.Use(middleware.AuthCOOKIEMiddleware())       //处理cookie验证
	r.Use(middleware.APIKeyAuthMiddleware())       //处理api密钥验证请求
	r.Use(middleware.CasbinMiddleware())           //处理casbin鉴权请求
	r.Use(middleware.CacheMiddleware())            //处理缓存请求
	r.Use(middleware.RateLimiterMiddleware())      //处理限流请求
	r.Use(middleware.MonitorMiddleware())          //处理流量监控请求
	r.Use(middleware.CSRFMiddleware())             //处理csrf防攻击请求
	// Route for user service
	//r.Any("/user/*proxyPath", func(c *gin.Context) {
	//	middleware.ProxyRequest(c, "user_service")
	//})
	//使用验证码限流
	r.GET("/api/captcha", middleware.RateLimiter(), func(c *gin.Context) {
		captcha := pkg.NewCaptcha()
		id, content, err := captcha.GenerateCaptcha()
		if err != nil {
			c.Error(gin.Error{
				Err:  err,
				Type: gin.ErrorTypePublic,
			})
		}
		data := map[string]string{
			"id": id, "content": content,
		}
		c.JSON(http.StatusOK, response.Success(data))
	})
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
