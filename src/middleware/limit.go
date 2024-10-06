package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

// 限流模块
var (
	redisClient *redis.Client
	ctx         = context.Background()
	rateLimit   = 100 // Max requests per window
	window      = time.Minute
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

// RateLimiterMiddleware implements a simple rate limiter using Redis
func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Unique key per IP
		ip := c.ClientIP()
		key := "rate_limiter_" + ip

		// Check current rate
		count, _ := redisClient.Get(ctx, key).Int()

		if count >= rateLimit {
			c.JSON(429, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		// Increment the count and set expiry
		redisClient.Incr(ctx, key)
		redisClient.Expire(ctx, key, window)

		c.Next()
	}
}

func RateLimiter(limit int, expire time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		key := "captcha_limit:" + clientIP

		// 使用 Redis 记录请求次数
		ctx := context.Background()
		count, _ := redisClient.Get(ctx, key).Int()

		if count >= limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "请求过多，请稍后重试"})
			return
		}

		// 增加计数并设置过期时间
		if count == 0 {
			redisClient.Set(ctx, key, 1, expire)
		} else {
			redisClient.Incr(ctx, key)
		}

		c.Next()
	}
}
