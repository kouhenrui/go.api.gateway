package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.api.gateway/src/config/response"
	"log"
	"net/http"
	"runtime/debug"
)

// ErrorHandler 全局错误捕获中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 捕获panic异常，并打印堆栈信息
				fmt.Println("捕获 panic:", r)
				debug.PrintStack()

				// 返回错误响应
				c.JSON(http.StatusInternalServerError, response.Error(
					http.StatusInternalServerError,
					"服务器内部错误",
				))
				c.Abort()
			}
		}()
		fmt.Println("全局错误捕捉处理中间件")
		c.Next() // 执行下一个中间件

		// 检查是否有错误
		//err := c.Errors.Last()
		//if err != nil {
		//	// 记录错误日志
		//	log.Printf("全局捕获错误: %v\n", err.Err)
		//
		//	// 返回统一的错误格式
		//	c.JSON(http.StatusInternalServerError, response.Error(http.StatusInternalServerError, "服务器内部错误"))
		//}
		// 捕获业务逻辑错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			log.Printf("业务错误: %v\n", err.Err)
			c.JSON(http.StatusInternalServerError, response.Error(http.StatusInternalServerError, err.Error()))
		}
	}
}
