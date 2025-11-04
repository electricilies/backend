package logger

import (
	"os"
	"time"

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

	timeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		loc, _ := time.LoadLocation(c.TimeZone)
		enc.AppendString(t.In(loc).Format(time.RFC3339))
	}
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
		encoderConfig.EncodeTime = timeEncoder
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
	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	Lgr = logger

	zap.ReplaceGlobals(Lgr)
}
