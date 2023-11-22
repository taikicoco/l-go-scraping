package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

var url string = "example.com"

func main() {
	c := colly.NewCollector()

	c.OnHTML("form", func(e *colly.HTMLElement) {
		formHtml, err := e.DOM.Html()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Form HTML:", formHtml)
	})

	err := c.Visit(url)
	if err != nil {
		panic(err)
	}
}
