package mtp

import "testing"

func TestHackerNewsMockImplementsInterfaces(t *testing.T) {
	implFeeder := func(HackerNewsFeeder) {}
	implScraper := func(HackerNewsScraper) {}
	implFeeder(&HackerNewsMock{})
	implScraper(&HackerNewsMock{})
}
