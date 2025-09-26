package logger

import (
	"backend/config"
	"path/filepath"
)

type loggerConfig struct {
	LogFile    string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	IsProdEnv  bool
}

func newDefaultLoggingConfig() *loggerConfig {
	return &loggerConfig{
		LogFile:    filepath.Join("logs", "app.log"),
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   false,
		IsProdEnv:  config.Cfg.EnvApp == "production",
	}
}

//TODO: read more config from file, env,...
