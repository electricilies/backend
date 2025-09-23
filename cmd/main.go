package main

import (
	"backend/config"
	"backend/internal/di"
)

func main() {
	config.LoadConfig()
	s := di.InitializeServer()
	s.Run()
}
