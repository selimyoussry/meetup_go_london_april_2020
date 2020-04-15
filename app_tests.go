package mtp

import (
	"fmt"
	"time"
)

const (
	testTimeout = 2 * time.Minute
)

// AppEndToEndTester -
type AppEndToEndTester struct {
	app *App

	// TestInterfaces
	HackerNewsFeeder HackerNewsFeeder
	DatabaseTester   DatabaseTester
}

// TestEndToEnd -
func (e *AppEndToEndTester) TestEndToEnd() error {
	// Make articles
	goArticle := Article{
		Title: "Go Debugging with Delve, or No More Fmt.Printfs",
		Link:  "https://tpaschalis.github.io/delve-debugging/",
	}
	unrelatedArticle := Article{
		Title: "Towards an ImageNet Moment for Speech-to-Text",
		Link:  "https://thegradient.pub/towards-an-imagenet-moment-for-speech-to-text/",
	}

	articles := []Article{
		goArticle, unrelatedArticle,
	}

	// Run the app in the background
	go e.app.Run()

	// Feed our fake hospital imaging database with new images
	err := e.HackerNewsFeeder.FeedArticles(articles)
	if err != nil {
		return err
	}

	// Wait for jobs on these images to happen
	tStart := time.Now()
	for {
		// List articles in our DB
		articles, err := e.DatabaseTester.GetAllArticles()
		if err != nil {
			return err
		}

		// Check the DB contains the right articles
		containsGoArticle := false
		containsUnrelatedArticle := false
		for _, article := range articles {
			if article.Link == goArticle.Link {
				containsGoArticle = true
			}
			if article.Link == unrelatedArticle.Link {
				containsUnrelatedArticle = true
			}
		}

		// If yes, return
		if containsGoArticle && !containsUnrelatedArticle {
			break
		}

		// Return an error after timeout
		if time.Now().Sub(tStart).Seconds() > testTimeout.Seconds() {
			return fmt.Errorf("End to end test timed out")
		}

		time.Sleep(time.Second)
	}

	// The test passes
	return nil
}
