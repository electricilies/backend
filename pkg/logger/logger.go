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
	if !c.EnableStdout && !c.EnableFile {
		Lgr = zap.NewNop()
		return
	}

	var cores []zapcore.Core
	logLevel := zap.NewAtomicLevelAt(zap.InfoLevel)
	if c.EnableStdout {
		stdout := zapcore.AddSync(os.Stdout)
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
		encoderConfig.TimeKey = "timestamp"
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		cores = append(cores, zapcore.NewCore(consoleEncoder, stdout, logLevel))
	}

	if c.EnableFile {
		file := zapcore.AddSync(&lumberjack.Logger{
			Filename:   c.LogFile,
			MaxSize:    c.MaxSize,
			MaxBackups: c.MaxBackups,
			MaxAge:     c.MaxAge,
			Compress:   c.Compress,
		})

		cfg := zap.NewProductionEncoderConfig()
		cfg.TimeKey = "timestamp"
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder

		fileEncoder := zapcore.NewJSONEncoder(cfg)
		cores = append(cores, zapcore.NewCore(fileEncoder, file, logLevel))
	}
	Lgr = zap.New(zapcore.NewTee(cores...))
}
