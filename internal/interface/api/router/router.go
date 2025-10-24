package router

import (
	"backend/internal/interface/api/handler"
	"backend/internal/interface/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Router interface {
	RegisterRoutes(engine *gin.Engine)
}

type router struct {
	userHandler       handler.User
	healthHandler     handler.HealthCheck
	metricMiddleware  middleware.Metric
	loggingMiddleware middleware.Logging
}

func NewRouter(userHandler handler.User, healthCheckHandler handler.HealthCheck, metricMiddleware middleware.Metric, loggingMiddleware middleware.Logging) Router {
	return &router{
		userHandler:       userHandler,
		healthHandler:     healthCheckHandler,
		metricMiddleware:  metricMiddleware,
		loggingMiddleware: loggingMiddleware,
	}
}

func (r *router) RegisterRoutes(engine *gin.Engine) {
	engine.Use(r.metricMiddleware.Handler())
	engine.Use(r.loggingMiddleware.Handler())
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	engine.GET("/health", r.healthHandler.Health)
	api := engine.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("", r.userHandler.List)
			users.POST("", r.userHandler.Create)
			users.GET("/:id", r.userHandler.Get)
			users.PUT("/:id", r.userHandler.Update)
			users.DELETE("/:id", r.userHandler.Delete)
		}
	}
}
