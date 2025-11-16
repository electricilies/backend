package server

import (
	"backend/internal/interface/api/handler"
	"backend/internal/interface/api/router"

	"github.com/gin-gonic/gin"
	swaggofiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine      *gin.Engine
	config      *Config
	authHandler handler.Auth // FIXME: Auth handler?
}

func ProvideServer(
	engine *gin.Engine,
	router router.Router,
	config *Config,
	authHandler handler.Auth,
) *Server {
	router.RegisterRoutes(engine)
	auth := engine.Group("/auth")
	{
		auth.GET("/*path", authHandler.Handler())
		auth.POST("/*path", authHandler.Handler())
	}
	engine.GET(
		"/swagger/*any",
		ginswagger.WrapHandler(
			swaggofiles.Handler,
			ginswagger.PersistAuthorization(true),
			ginswagger.Oauth2DefaultClientID("swagger"),
		),
	)
	return &Server{
		engine:      engine,
		config:      config,
		authHandler: authHandler,
	}
}

func (s *Server) Run() error {
	return s.engine.Run()
}
