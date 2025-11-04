package main

import (
	"backend/config"
	"backend/internal/di"
	"backend/pkg/logger"
	"io"
	"log"

	_ "backend/docs"

	"github.com/gin-gonic/gin"
)

// @BasePath								/api
//
// @securitydefinitions.oauth2.accessCode	OAuth2AccessCode
//
// @tokenUrl								auth/realms/electricilies/protocol/openid-connect/token
// @authorizationurl						auth/realms/electricilies/protocol/openid-connect/auth
// @scope.write							Grants write access
func main() {
	config.LoadConfig()
	s := di.InitializeServer()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	logger.InitializeLogger()
	err := s.Run()
	if err != nil {
		log.Fatal("Server run error", err)
	}
}
