package main

import (
	"go.uber.org/zap"
)

func main() {
	// Create a sugared logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // Flushes buffer, if any
	sugar := logger.Sugar()

	// Log a simple message
	sugar.Info("This is a sugared logger example")

	// Log a message with fields
	sugar.Infow("A sugared logger example with fields",
		"field1", "value1",
		"field2", 42,
	)
}