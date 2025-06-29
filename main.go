package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	base_url := args[0]

	const maxConcurrency = 4
	config, err := configure(base_url, maxConcurrency)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s ...\n", base_url)

	start := time.Now()

	config.wg.Add(1)
	go config.crawlPage(base_url)
	config.wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Crawler took %v to complete!\n", elapsed)

	for normalizedURL, visitCount := range config.pages {
		fmt.Printf("%d - %s\n", visitCount, normalizedURL)
	}
}
