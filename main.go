package main

import (
	"fmt"
)

var url string = "example.com"

func main() {
	getByGoquery(url)
	inputFields := getByGoquery(url)

	for _, field := range inputFields {
		fmt.Printf("%+v\n", field)
	}
}
