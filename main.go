package main

import (
	"flag"
	"fmt"
	"log"
	"server/scraping"
)

func main() {
	urlPtr := flag.String("url", "exsample.com", "URL to process")
	flag.Parse()
	url := *urlPtr

	if url == "" {
		log.Fatal("URL environment variable is not set")
	}

	inputFields := scraping.GetByGoquery(url)
	for _, field := range inputFields {
		fmt.Printf("%+v\n", field)
	}
}
