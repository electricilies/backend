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
}

func newDefaultLoggingConfig() *loggerConfig {
	return &loggerConfig{
		LogFile:      filepath.Join("logs", "app.log"),
		MaxSize:      10,
		MaxBackups:   3,
		MaxAge:       7,
		Compress:     false,
		EnableStdout: config.Cfg.EnableStdout,
		EnableFile:   config.Cfg.EnableFile,
	}
}

// TODO: read more config from file, env,...
