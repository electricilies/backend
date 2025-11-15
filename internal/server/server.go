package server

import (
	"backend/config"
	"backend/internal/interface/api/handler"
	"backend/internal/interface/api/router"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine      *gin.Engine
	srvCfg      *config.Server
	authHandler handler.Auth
}

func New(e *gin.Engine, r router.Router, srvCfg *config.Server, authHandler handler.Auth) *Server {
	r.RegisterRoutes(e)
	auth := e.Group("/auth")
	{
		auth.GET("/*path", authHandler.Handler())
		auth.POST("/*path", authHandler.Handler())
	}
	e.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerfiles.Handler,
			ginSwagger.PersistAuthorization(true),
			ginSwagger.Oauth2DefaultClientID("swagger"),
		),
	)
	return &Server{
		engine:      e,
		srvCfg:      srvCfg,
		authHandler: authHandler,
	}
}

func (s *Server) Run() error {
	return s.engine.Run()
}
