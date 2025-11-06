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

type loggingMiddleware struct {
	skipPaths []string
}

func NewLogging() Logging {
	return &loggingMiddleware{
		skipPaths: []string{"/health", "/metrics", "/swagger"},
	}
}

func (l *loggingMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		skipPathSet := make(map[string]struct{}, len(l.skipPaths))
		for _, p := range l.skipPaths {
			skipPathSet[p] = struct{}{}
		}
		if _, ok := skipPathSet[ctx.Request.URL.Path]; ok {
			ctx.Next()
			return
		}
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
				logger.Lgr.Error(e, fields...)
			}
			return
		}
		logger.Lgr.Info(path, fields...)
	}
}
