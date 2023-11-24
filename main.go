package main

import (
	"fmt"
	"log"
	"os"
	"server/scraping"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var url string = os.Getenv("URL")
	if url == "" {
		log.Fatal("URL environment variable is not set")
	}

	inputFields := scraping.GetByGoquery(url)
	for _, field := range inputFields {
		fmt.Printf("%+v\n", field)
	}
}
