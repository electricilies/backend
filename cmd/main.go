package main

import (
	"io"
	"log"

	"backend/config"
	"backend/internal/di"
	"backend/pkg/logger"

	"github.com/gin-gonic/gin"
)

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
