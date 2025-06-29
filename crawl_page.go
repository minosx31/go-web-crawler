package main

import (
	"fmt"
	"net/url"
)

// func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{} // Insert to channel
	defer func() {
		<-cfg.concurrencyControl // Remove from channel
		cfg.wg.Done()            // Decrement WaitGroup's counter
	}()

	if cfg.hasVisitedMaxPages() {
		return
	}

	current, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	// Compare both domains, skip if different
	if current.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	// Obtain normalized URL to use as key
	normCurrentUrl, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
		return
	}

	// If url has been visited, return
	isFirstVisit := cfg.addPageVisit(normCurrentUrl)
	if !isFirstVisit {
		return
	}

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}
	fmt.Printf("Crawling %s: %s\n", rawCurrentURL, htmlBody)

	nextURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
		return
	}

	fmt.Printf("Number of URLs found: %d\n", len(nextURLs))

	for _, nextURL := range nextURLs {
		cfg.wg.Add(1) // Increment WaitGroup's counter
		go cfg.crawlPage(nextURL)
	}
}
