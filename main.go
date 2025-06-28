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
	fmt.Printf("starting crawl of: %s...", base_url)

	htmlBody, err := getHTML(base_url)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(htmlBody)
}
