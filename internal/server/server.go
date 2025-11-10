package server

import (
	"net/http"
	"strings"

	"backend/config"
	"backend/internal/interface/api/router"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine *gin.Engine
	cfg    *config.Config
}

func New(e *gin.Engine, r router.Router, cfg *config.Config) *Server {
	r.RegisterRoutes(e)
	auth := e.Group("/auth")
	{
		handler := authHandler(cfg)
		auth.GET("/*path", handler)
		auth.POST("/*path", handler)
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
		engine: e,
		cfg:    cfg,
	}
}

func (s *Server) Run() error {
	return s.engine.Run()
}

func authHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		basePath := cfg.KcBasePath
		redirectURL := basePath + strings.TrimPrefix(c.Request.URL.String(), "/auth")
		c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	}
}
