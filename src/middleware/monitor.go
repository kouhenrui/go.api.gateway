package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 流量监控
var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_gateway_requests_total",
			Help: "Total number of requests handled by the API Gateway",
		},
		[]string{"method", "endpoint"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_gateway_request_duration_seconds",
			Help:    "Duration of requests handled by the API Gateway",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	// Register the metrics with Prometheus
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)
}

// MonitorMiddleware monitors request metrics for Prometheus
func MonitorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		timer := prometheus.NewTimer(requestDuration.WithLabelValues(c.Request.Method, c.FullPath()))
		requestCount.WithLabelValues(c.Request.Method, c.FullPath()).Inc()

		c.Next()

		timer.ObserveDuration()
	}
}

// ExposePrometheusMetrics exposes the Prometheus metrics endpoint
func ExposePrometheusMetrics() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}
