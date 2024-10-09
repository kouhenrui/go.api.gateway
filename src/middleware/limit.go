package middleware

import (
	"github.com/gin-gonic/gin"
	"go.api.gateway/src/config"
	"go.api.gateway/src/config/response"
	"net/http"
	"strconv"
	"time"
)

// 限流模块
var (
	redisClient config.RedisClient
	rateLimit   = 100 // Max requests per window
	window      = time.Minute
)

//func init() {
//	redisClient = redis.NewClient(&redis.Options{
//		Addr: "localhost:6379",
//	})
//}

// RateLimiterMiddleware implements a simple rate limiter using Redis
func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Unique key per IP
		ip := c.ClientIP()
		key := "rate_limiter_" + ip

		// Check current rate
		count, _ := redisClient.Get(key)
		counts, _ := strconv.Atoi(count)
		if counts >= rateLimit {
			c.JSON(http.StatusTooManyRequests, response.Error(http.StatusTooManyRequests, "请求过多，请稍后重试")) //gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		// Increment the count and set expiry
		redisClient.Incr(key)

		redisClient.Expire(key, window)

		c.Next()
	}
}

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		key := "captcha_limit:" + clientIP

		count, _ := redisClient.Get(key)
		counts, _ := strconv.Atoi(count)
		if counts >= config.ViperConfig.Captcha.Limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, response.Error(http.StatusTooManyRequests, "请求过多，请稍后重试")) // gin.H{"message": "请求过多，请稍后重试"})
			return
		}

		// 增加计数并设置过期时间
		if counts == 0 {
			redisClient.Set(key, 1, config.ViperConfig.Captcha.Expired)
		} else {
			redisClient.Incr(key)
		}

		c.Next()
	}
}
