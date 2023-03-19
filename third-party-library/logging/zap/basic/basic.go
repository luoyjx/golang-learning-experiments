package main

// Import necessary packages
import (
	"go.uber.org/zap"
)

func main() {
	// Create a logger with default settings
	logger, _ := zap.NewProduction()
	defer logger.Sync() // Flush any buffered log entries

	// Log a simple message at different log levels
	logger.Debug("Zap logger debug level")
	logger.Info("Zap logger info level")
	logger.Warn("Zap logger warn level")
	logger.Error("Zap logger error level")

	// Log a message with fields at different log levels
	logger.Debug("Zap logger with fields at debug level", zap.String("key", "value"), zap.Int("number", 42))
	logger.Info("Zap logger with fields at info level", zap.String("key", "value"), zap.Int("number", 42))
	logger.Warn("Zap logger with fields at warn level", zap.String("key", "value"), zap.Int("number", 42))
	logger.Error("Zap logger with fields at error level", zap.String("key", "value"), zap.Int("number", 42))
}
