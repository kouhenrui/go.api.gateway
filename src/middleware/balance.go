package middleware

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	// Store round robin counters for each service
	roundRobinCounters = make(map[string]int)
	mu                 sync.Mutex
)

// 负载均衡模块

// ProxyRequest handles the request routing and load balancing
func ProxyRequest(c *gin.Context, serviceName string) {
	services := []string{
		"http://localhost:8081",
		"http://localhost:8082",
	}

	// Get the next service endpoint using round robin
	endpoint := getNextService(services)

	proxyURL := endpoint + c.Param("proxyPath")
	proxyReq, _ := http.NewRequest(c.Request.Method, proxyURL, c.Request.Body)

	// Copy headers from the original request
	for key, values := range c.Request.Header {
		for _, value := range values {
			proxyReq.Header.Add(key, value)
		}
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(proxyReq)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Service unavailable"})
		return
	}
	defer resp.Body.Close()

	// Copy headers and status code from response
	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}
	c.Status(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}

// getNextService returns the next service endpoint using round-robin algorithm
func getNextService(services []string) string {
	mu.Lock()
	defer mu.Unlock()

	counter := roundRobinCounters["user_service"]
	endpoint := services[counter%len(services)]
	roundRobinCounters["user_service"] = counter + 1
	return endpoint
}
