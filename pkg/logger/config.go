package logger

import (
	"path/filepath"

	"backend/config"
)

type loggerConfig struct {
	LogFile      string
	MaxSize      int
	MaxBackups   int
	MaxAge       int
	Compress     bool
	EnableFile   bool
	EnableStdout bool
	TimeZone     string
}

func NewConfig(cfg *config.Config) *loggerConfig {
	return &loggerConfig{
		LogFile:      filepath.Join("logs", "app.log"),
		MaxSize:      10,
		MaxBackups:   3,
		MaxAge:       7,
		Compress:     false,
		EnableStdout: cfg.EnableStdout,
		EnableFile:   cfg.EnableFile,
		TimeZone:     cfg.TimeZone,
	}
}

// TODO: read more config from file, env,...
