package router

import (
	handler "backend/internal/interface/api/handler"
	"backend/internal/interface/api/middleware"

	"github.com/gin-gonic/gin"
)

type Router interface {
	RegisterRoutes(engine *gin.Engine)
}

type router struct {
	userHandler      handler.User
	healthHanler     handler.HealthCheck
	metricMiddleware middleware.Metric
}

func NewRouter(userHandler handler.User, healthCheckHandler handler.HealthCheck, metricMiddleware middleware.Metric) Router {
	return &router{
		userHandler:      userHandler,
		healthHanler:     healthCheckHandler,
		metricMiddleware: metricMiddleware,
	}
}

func (r *router) RegisterRoutes(engine *gin.Engine) {
	engine.GET("health", r.healthHanler.Get)
	engine.Use(r.metricMiddleware.Handler())
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
