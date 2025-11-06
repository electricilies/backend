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
// @securitydefinitions.oauth2.accessCode	OAuth2AccessCode
//
// @tokenUrl								/auth/realms/electricilies/protocol/openid-connect/token
// @authorizationUrl						/auth/realms/electricilies/protocol/openid-connect/auth
//
// @securitydefinitions.oauth2.password	OAuth2Password
//
// @tokenUrl								/auth/realms/electricilies/protocol/openid-connect/token
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
