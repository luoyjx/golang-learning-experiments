// Import necessary packages
import (
	"go.uber.org/zap"
)

func main() {
	// Create a logger with default settings
	logger, _ := zap.NewProduction()
	defer logger.Sync() // Flush any buffered log entries

	// Log a simple message
	logger.Info("Zap logger basic usage example")

	// Log a message with fields
	logger.Info("Zap logger with fields", zap.String("key", "value"), zap.Int("number", 42))
}