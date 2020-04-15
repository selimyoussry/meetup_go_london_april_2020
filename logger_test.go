package mtp

import (
	"testing"

	"go.uber.org/zap"
)

func TestLoggerImpl(t *testing.T) {
	implLogger := func(Logger) {}
	logger, err := zap.NewProduction()
	if err != nil {
		t.Fatalf("Could not create Zap Logger, err=%s", err)
	}
	implLogger(logger.Sugar())
}
