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
}

func New(e *gin.Engine, r router.Router) *Server {
	r.RegisterRoutes(e)
	auth := e.Group("/auth")
	{
		auth.GET("/*path", authHandler)
		auth.POST("/*path", authHandler)
	}
	e.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerfiles.Handler,
			ginSwagger.PersistAuthorization(true),
			ginSwagger.Oauth2DefaultClientID("frontend"),
		),
	)
	return &Server{
		engine: e,
	}
}

func (s *Server) Run() error {
	return s.engine.Run()
}

func authHandler(c *gin.Context) {
	redirectURL := config.Cfg.KcBasePath + strings.TrimPrefix(c.Request.URL.String(), "/auth")
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
