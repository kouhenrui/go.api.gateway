package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

// 缓存模块
func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

// CacheMiddleware caches the GET requests using Redis
func CacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			c.Next()
			return
		}

		cacheKey := generateCacheKey(c.Request)

		cachedResponse, err := redisClient.Get(ctx, cacheKey).Result()
		if err == redis.Nil {
			// Cache miss, continue to the handler and cache the response
			c.Next()

			// Cache the response
			//if c.Writer.Status() == http.StatusOK {
			//	//responseBody, _ := ioutil.ReadAll(c.Writer.Body)
			//	//redisClient.Set(ctx, cacheKey, responseBody, 5*time.Minute)
			//}
		} else if err == nil {
			// Cache hit, return cached response
			c.Writer.Write([]byte(cachedResponse))
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// generateCacheKey generates a unique cache key based on the URL and request parameters
func generateCacheKey(req *http.Request) string {
	hash := sha256.New()
	hash.Write([]byte(req.URL.String() + req.URL.RawQuery))
	return hex.EncodeToString(hash.Sum(nil))
}
