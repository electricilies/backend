package logger

import (
	"backend/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitializeLogger() {
	if config.Cfg.EnvApp == "production" {
		Logger = newProductionLogger()
		return
	}
	Logger = newDevelopmentLogger()
}

func newDevelopmentLogger() *zap.Logger {
	stdout := zapcore.AddSync(os.Stdout)
	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(cfg)
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
	)
	return zap.New(core)
}

func newProductionLogger() *zap.Logger {
	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log", //TODO: add config instead of hard code
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     7,
	})
	level := zap.NewAtomicLevelAt(zap.InfoLevel)

	cfg := zap.NewProductionEncoderConfig()
	cfg.TimeKey = "timestamp"
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(cfg)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, file, level),
	)

	return zap.New(core)
}
