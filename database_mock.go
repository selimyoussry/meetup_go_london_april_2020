package mtp

import (
	"sync"
)

// DatabaseMock contains a list of articles
type DatabaseMock struct {
	articlesLock sync.RWMutex
	articles     map[string]Article

	waiter Waiter
	failer Failer

	logger Logger
}

// GetAllArticles gets the homepage of hacker news
func (s *DatabaseMock) GetAllArticles() ([]Article, error) {
	s.waiter.Wait()
	err := s.failer.Fails()
	if err != nil {
		s.logger.Errorf("GetAllArticles failed with err=%s", err)
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

// StoreArticles stores articles in the mock HackerNews map
// here for testing purposes
func (s *DatabaseMock) StoreArticles(articles []Article) error {
	s.waiter.Wait()
	err := s.failer.Fails()
	if err != nil {
		s.logger.Errorf("StoreArticles failed with err=%s", err)
		return err
	}

	s.articlesLock.Lock()
	defer s.articlesLock.Unlock()
	for _, article := range articles {
		if !isGoArticle(article) {
			continue
		}
		s.articles[article.Link] = article
	}
	return nil
}
