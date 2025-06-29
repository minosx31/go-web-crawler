package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// usage: ./crawler URL maxConcurrency maxPages
func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		os.Exit(1)
	} else if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	base_url := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("error parsing argument to int")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("error parsing argument to int")
		os.Exit(1)
	}

	config, err := configure(base_url, maxConcurrency, maxPages)
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

	printReport(config.pages, config.baseURL.String())
}
