package scraping

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type InputField struct {
	Tag  string
	Name string
	Type string
}

func GetByGoquery(url string) []InputField {
	var fields []InputField
	webPage := (url)
	res, err := http.Get(webPage)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	htmlTags := []string{"input"}

	for _, tag := range htmlTags {
		doc.Find(tag).Each(func(i int, s *goquery.Selection) {
			field := InputField{
				Tag: tag,
			}

			field.Name, _ = s.Attr("name")
			field.Type, _ = s.Attr("type")

			fields = append(fields, field)
		})
	}
	return fields
}
