package mtp

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestAppMockEndToEnd(t *testing.T) {
	// Setup
	waiter := &WaitFixed{WaitDuration: time.Second}
	failer := &FailRandom{FailRate: 0.05}
	zapLogger, err := zap.NewProduction()
	if err != nil {
		t.Fatalf("Could not create Zap logger, err=%s", err)
	}
	logger := zapLogger.Sugar()

	hackerNews := &HackerNewsMock{
		articles: map[string]Article{},
		waiter:   waiter,
		failer:   failer,
		logger:   logger,
	}
	database := &DatabaseMock{
		articles: map[string]Article{},
		waiter:   waiter,
		failer:   failer,
		logger:   logger,
	}
	app := &App{
		hackerNewsScraper: hackerNews,
		databaseConnecter: database,
		logger:            logger,
	}

	// Run the test
	tester := &AppEndToEndTester{
		app:              app,
		HackerNewsFeeder: hackerNews,
		DatabaseTester:   database,
	}
	err = tester.TestEndToEnd()
	if err != nil {
		t.Fatalf("TestEndToEnd: Expected err=nil, got err=%s", err)
	}
}
