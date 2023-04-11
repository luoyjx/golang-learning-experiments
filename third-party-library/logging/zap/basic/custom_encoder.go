package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Create a custom encoder configuration
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Create a custom encoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// Create a console output
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	// Create a core
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, consoleErrors, zapcore.ErrorLevel),
		zapcore.NewCore(encoder, consoleDebugging, zapcore.DebugLevel),
	)

	// Create a logger
	logger := zap.New(core)

	// Use the logger
	logger.Info("Hello, world!")
}
