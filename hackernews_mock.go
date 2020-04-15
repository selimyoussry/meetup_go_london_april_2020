package mtp

import (
	"sync"
)

// HackerNewsMock contains a list of articles
type HackerNewsMock struct {
	articlesLock sync.RWMutex
	articles     map[string]Article

	waiter Waiter
	failer Failer

	logger Logger
}

// GetLatestArticles gets the homepage of hacker news
func (s *HackerNewsMock) GetLatestArticles() ([]Article, error) {
	s.waiter.Wait()
	err := s.failer.Fails()
	if err != nil {
		s.logger.Errorf("GetLatestArticles failed with err=%s", err)
		return nil, err
	}

	var articles []Article

	s.articlesLock.RLock()
	defer s.articlesLock.RUnlock()

	// Get all articles in the map
	for _, article := range s.articles {
		articles = append(articles, article)
	}

	return articles, nil
}

// FeedArticles stores articles in the mock HackerNews map
// here for testing purposes
func (s *HackerNewsMock) FeedArticles(articles []Article) error {
	s.waiter.Wait()
	err := s.failer.Fails()
	if err != nil {
		s.logger.Errorf("FeedArticles failed with err=%s", err)
		return err
	}

	s.articlesLock.Lock()
	defer s.articlesLock.Unlock()
	for _, article := range articles {
		s.articles[article.Link] = article
	}

	s.logger.Infof("Fed %d articles", len(articles))
	return nil
}
