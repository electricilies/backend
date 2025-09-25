package main

import (
	"backend/config"
	"backend/internal/di"
	"backend/pkg/logger"
)

func main() {
	config.LoadConfig()
	s := di.InitializeServer()
	logger.InitializeLogger()
	s.Run()
}
