package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Lgr *zap.Logger

func InitializeLogger() {
	c := newDefaultLoggingConfig()
	if c.IsProdEnv {
		Lgr = newProductionLogger(c)
	}
	Lgr = newDevelopmentLogger()
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

func newProductionLogger(config *loggerConfig) *zap.Logger {
	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.LogFile,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	})
	level := zap.NewAtomicLevelAt(zap.InfoLevel)

	cfg := zap.NewProductionEncoderConfig()
	cfg.TimeKey = "timestamp"
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(cfg)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, file, level),
	)

	return zap.New(core, zap.AddCaller())
}
