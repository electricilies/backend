package main

import (
	"backend/config"
	"backend/internal/di"
	"backend/pkg/logger"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	s := di.InitializeServer()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	logger.InitializeLogger()
	s.Run()
}
