package server

import (
	"backend/docs"
	"backend/internal/interface/api/router"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(e *gin.Engine, r router.Router) *Server {

	docs.SwaggerInfo.BasePath = "api/v1"
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.RegisterRoutes(e)

	return &Server{
		engine: e,
	}
}

func (s Server) Run() error {
	return s.engine.Run(":8080")
}
