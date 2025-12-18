package logger

import "go.uber.org/zap"

// NewLogger creates a new zap.Logger with sensible defaults.
func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return logger
}


