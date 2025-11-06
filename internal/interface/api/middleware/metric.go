package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metric interface {
	Handler() gin.HandlerFunc
}

type metricMiddleware struct {
	httpRequestsTotal   *prometheus.CounterVec
	httpRequestDuration *prometheus.HistogramVec
}

func NewMetric() Metric {
	return &metricMiddleware{
		httpRequestsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"method", "path", "status"},
		),
		httpRequestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "HTTP request duration in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"method", "path", "status"},
		),
	}
}

func (m *metricMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		duration := time.Since(start).Seconds()
		status := strconv.Itoa(ctx.Writer.Status())
		path := ctx.FullPath()
		if path == "" {
			path = ctx.Request.URL.Path
		}

		m.httpRequestsTotal.WithLabelValues(ctx.Request.Method, path, status).Inc()
		m.httpRequestDuration.WithLabelValues(ctx.Request.Method, path, status).Observe(duration)
	}
}
