package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// Parse both rawBaseURL and rawCurrentURL
	base, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	current, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}

	// Compare both domains, skip if different
	if base.Hostname() != current.Hostname() {
		return
	}

	// Obtain normalized URL to use as key
	normCurrentUrl, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
		return
	}

	// If visited, increment count and return
	if _, exists := pages[normCurrentUrl]; exists {
		pages[normCurrentUrl]++
		return
	}

	// Mark as visited
	pages[normCurrentUrl] = 1

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}
	fmt.Printf("Crawling %s: %s\n", rawCurrentURL, htmlBody)

	toCrawlUrls, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
		return
	}

	fmt.Printf("Number of URLs found: %d\n", len(toCrawlUrls))

	for _, url := range toCrawlUrls {
		crawlPage(rawBaseURL, url, pages)
	}
}
