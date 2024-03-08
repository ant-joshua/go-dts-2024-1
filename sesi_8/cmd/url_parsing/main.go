package main

import (
	"fmt"
	"net/url"
)

func main() {
	var searchUrl = "https://www.google.com/search?q=golang&language=en"

	result, err := url.Parse(searchUrl)

	if err != nil {
		println(err.Error())
		return
	}

	println("Scheme  :", result.Scheme)
	println("Host    :", result.Host)
	println("Path    :", result.Path)
	println("RawQuery:", result.RawQuery)

	println("Query   :", result.Query().Get("q"))
	fmt.Println("Language:", result.Query().Has("language"))

}
