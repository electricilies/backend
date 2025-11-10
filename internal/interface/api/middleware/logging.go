package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logging interface {
	Handler() gin.HandlerFunc
}

type loggingMiddleware struct {
	logger *zap.Logger
}

func NewLogging(logger *zap.Logger) Logging {
	return &loggingMiddleware{
		logger: logger,
	}
}

func (l *loggingMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		ctx.Next()
		end := time.Now()
		latency := end.Sub(start)

		status := ctx.Writer.Status()
		fields := []zapcore.Field{
			zap.Int("status", status),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.Duration("latency", latency),
		}

		if len(ctx.Errors) > 0 && status >= 500 {
			for _, e := range ctx.Errors.Errors() {
				l.logger.Error(e, fields...)
			}
			return
		}
		l.logger.Info(path, fields...)
	}
}
