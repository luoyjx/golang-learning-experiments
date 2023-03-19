package main

// Import necessary packages
import (
	"go.uber.org/zap"
	"os"
)

func main() {
	// Create a local file output writer
	file, err := os.Create("logs.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Configure the custom zap logger
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// Add the local file output writer to the logger
	config.OutputPaths = append(config.OutputPaths, file.Name())

	// Build the custom zap logger
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// Use the custom zap logger
	logger.Info("This is a custom zap logger with a local file output writer.")
}
