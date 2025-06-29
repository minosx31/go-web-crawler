package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody string, rawBaseURL *url.URL) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	nodes, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	// Traverse node tree to get all urls enclosed in anchor tags
	var urls []string
	for n := range nodes.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href, err := url.Parse(a.Val)
					if err != nil {
						fmt.Printf("couldn't parse href '%v': %v\n", a.Val, err)
						continue
					}

					resolvedURL := rawBaseURL.ResolveReference(href)
					urls = append(urls, resolvedURL.String())
				}
			}
		}
	}

	return urls, nil
}
