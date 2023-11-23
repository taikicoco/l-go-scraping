package scraping

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func getByColly(url string) {
	c := colly.NewCollector()
	var forms [][]string
	c.OnHTML("form", func(e *colly.HTMLElement) {
		formHtml, err := e.DOM.Html()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		form := getFormFields(formHtml)
		forms = append(forms, form)
	})
	err := c.Visit(url)
	if err != nil {
		panic(err)
	}
	for _, form := range forms {
		for _, field := range form {
			fmt.Println(field)
		}
	}
}

func getFormFields(html string) []string {
	var elements []string
	inputTags := []string{"<input", "<select", "<textarea", "<label", "<option", "<a"}
	start := -1

	for i := 0; i < len(html); i++ {
		if start == -1 {
			for _, tag := range inputTags {
				if strings.HasPrefix(html[i:], tag) {
					start = i
					break
				}
			}
		} else if html[i] == '>' {
			elements = append(elements, html[start:i+1])
			start = -1
		}
	}
	return elements
}

func getFormFieldsWithTag(html string) []string {
	var elements []string
	start := -1

	for i := 0; i < len(html); i++ {
		if html[i] == '<' {
			start = i
		}
		if html[i] == '>' && start != -1 {
			elements = append(elements, html[start:i+1])
			start = -1
		}
	}
	return elements
}
