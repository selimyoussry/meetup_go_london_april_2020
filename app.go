package mtp

import "time"

/*
  Side structs to make it work
*/

// Article contains a HackerNews's title and link
type Article struct {
	Title string
	Link  string
}

/*
  Service clients interfaces
*/

// HackerNewsScraper implements what our HackerNews scraper should be able to do
type HackerNewsScraper interface {
	GetLatestArticles() ([]Article, error)
}

// DatabaseConnecter implements the interactions our app has with our internal database
type DatabaseConnecter interface {
	StoreArticles([]Article) error
}

/*
  Testing interfaces
*/

// HackerNewsFeeder can feed data to our fake HackerNews website
type HackerNewsFeeder interface {
	FeedArticles([]Article) error
}

// DatabaseTester implements the methods to check what's in the DB
type DatabaseTester interface {
	GetAllArticles() ([]Article, error)
}

/*
  Convenience interfaces
*/

// Logger implements your expected logging functions
type Logger interface {
	Debugf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}

/*
  App
*/

// App implements our business logic, and can interact with external services
type App struct {
	// External services the app can interact with
	hackerNewsScraper HackerNewsScraper
	databaseConnecter DatabaseConnecter

	// Logger
	logger Logger
}

/*
  App Business Logic
*/

// Run -
func (a *App) Run() error {
	for {
		// Get latest articles
		latestArticles, err := a.hackerNewsScraper.GetLatestArticles()
		if err != nil {
			return err
		}
		a.logger.Infof("GetLatestArticles returned %d articles", len(latestArticles))

		// Store in DB
		err = a.databaseConnecter.StoreArticles(latestArticles)
		if err != nil {
			return err
		}
		a.logger.Infof("Stored %d articles", len(latestArticles))

		time.Sleep(defaultScapeLoopSleepDuration)
	}
}
