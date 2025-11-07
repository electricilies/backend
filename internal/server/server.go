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
	basePath := config.Cfg.KcBasePath
	if env := config.Cfg.SwaggerEnv; env != "" {
		basePath = strings.Replace(basePath, "keycloak", "keycloak."+env, 1)
	}
	redirectURL := basePath + strings.TrimPrefix(c.Request.URL.String(), "/auth")
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
