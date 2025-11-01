package server

import (
	"backend/internal/interface/api/router"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine *gin.Engine
}

func New(e *gin.Engine, r router.Router) *Server {
	r.RegisterRoutes(e)
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return &Server{
		engine: e,
	}
}

func (s *Server) Run() error {
	return s.engine.Run()
}
