package main

import (
	"fmt"
	"web-crawler/web_crawler"
)

func main() {
	entries := []string{""}
	wc := web_crawler.New(entries)

	fmt.Printf("Got the web crawler: %v", wc)
}
