package main

import (
	"fmt"
	"net/url"
	"sync"
)

// Config struct for goroutines to share access
type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	_, exists := cfg.pages[normalizedURL]
	isFirst = !exists

	// Critical section
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if exists {
		cfg.pages[normalizedURL]++
	} else {
		cfg.pages[normalizedURL] = 1
	}
	return
}

func configure(rawBaseURL string, maxConcurrency int, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL %s: %v", rawBaseURL, err)
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}, nil
}

func (cfg *config) hasVisitedMaxPages() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages) >= cfg.maxPages
}
