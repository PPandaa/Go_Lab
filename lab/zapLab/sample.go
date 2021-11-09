package zapLab

import (
	"time"

	"go.uber.org/zap"
)

func CommonCase() {

	Logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

}

func SuggarCase() {

	Logger.Sugar().Infow("failed to fetch URL",
		"attempt", 3,
		"backoff", time.Second,
	)
	Logger.Sugar().Infof("Failed to fetch")

}
