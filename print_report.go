package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	url        string
	visitCount int
}

func printReport(pages map[string]int, baseURL string) {
	output := ""
	separator := "========================================================"
	header := fmt.Sprintf("REPORT for %s", baseURL)
	output = fmt.Sprintf("%s\n%s\n%s\n", separator, header, separator)

	sortedPages := sortPages(pages)

	for _, entry := range sortedPages {
		output += fmt.Sprintf("Found %d internal links to %s\n", entry.visitCount, entry.url)
	}

	fmt.Println(output)
}

func sortPages(pages map[string]int) []Entry {
	// Convert map to slice
	var entries []Entry
	for url, visitCount := range pages {
		entries = append(entries, Entry{
			url:        url,
			visitCount: visitCount,
		})
	}

	// Sort slice
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].visitCount != entries[j].visitCount {
			return entries[i].visitCount > entries[j].visitCount
		}
		return entries[i].url < entries[j].url
	})

	return entries
}
