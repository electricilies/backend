package middleware

import (
	"backend/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logging interface {
	Handler() gin.HandlerFunc
}

type logging struct {
	skipPaths []string
}

func NewLogging() Logging {
	return &logging{
		skipPaths: []string{"/health", "/metrics", "/swagger"},
	}
}

func (l *logging) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		skipPathSet := make(map[string]struct{}, len(l.skipPaths))
		for _, p := range l.skipPaths {
			skipPathSet[p] = struct{}{}
		}
		if _, ok := skipPathSet[c.Request.URL.Path]; ok {
			c.Next()
			return
		}
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		status := c.Writer.Status()
		fields := []zapcore.Field{
			zap.Int("status", status),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		}

		if len(c.Errors) > 0 && status >= 500 {
			for _, e := range c.Errors.Errors() {
				logger.Lgr.Error(e, fields...)
			}
			return
		}
		logger.Lgr.Info(path, fields...)
	}
}
