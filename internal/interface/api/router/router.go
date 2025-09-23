package router

import (
	handler "backend/internal/interface/api/handler"

	"github.com/gin-gonic/gin"
)

type Router interface {
	RegisterRoutes(engine *gin.Engine)
}

type router struct {
	userHandler handler.User
}

func NewRouter(userHandler handler.User) Router {
	return &router{
		userHandler: userHandler,
	}
}

func (r *router) RegisterRoutes(engine *gin.Engine) {
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
