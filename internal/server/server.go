package server

import (
	"backend/config"
	"backend/internal/interface/api/router"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine *gin.Engine
}

func New(e *gin.Engine, r router.Router) *Server {
	r.RegisterRoutes(e)
	e.GET("/auth/*path", func(c *gin.Context) {
		redirectURL := config.Cfg.KcBasePath + strings.TrimPrefix(c.Request.URL.String(), "/auth")
		c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	})
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return &Server{
		engine: e,
	}
}

func (s *Server) Run() error {
	return s.engine.Run()
}
