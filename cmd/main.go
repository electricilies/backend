package main

import (
	"backend/config"
	"backend/internal/di"
	"backend/pkg/logger"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

// @BasePath								/api
// @securitydefinitions.apikey	Admin
// @in							header
// @name						Authorization
// @securitydefinitions.apikey Customer
// @in							header
// @name						Authorization
// @securitydefinitions.apikey Staff
// @in							header
// @name						Authorization

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
