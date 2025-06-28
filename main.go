package main

import (
	"fmt"
	"os"
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
	fmt.Printf("starting crawl of: %s ...\n", base_url)

	pages := make(map[string]int)
	crawlPage(base_url, base_url, pages)
}
