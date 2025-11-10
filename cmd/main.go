package main

import (
	"io"
	"log"

	"backend/config"
	"backend/internal/di"
	"backend/pkg/logger"

	_ "backend/docs"

	"github.com/gin-gonic/gin"
)

// @BasePath								/api
//
// @securitydefinitions.oauth2.password	OAuth2Password
// @tokenUrl								/auth/realms/electricilies/protocol/openid-connect/token
//
// @securitydefinitions.oauth2.accessCode	OAuth2AccessCode
// @tokenUrl								/auth/realms/electricilies/protocol/openid-connect/token
// @authorizationUrl						/auth/realms/electricilies/protocol/openid-connect/auth
func main() {
	config.NewDefaultConfig()
	s := di.InitializeServer()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	logger.New()
	err := s.Run()
	if err != nil {
		log.Fatal("Server run error", err)
	}
}
