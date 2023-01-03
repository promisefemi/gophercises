package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/promisefemi/gophercises/link"
)

func main() {
	urlFlag := flag.String("url", "http://gophercises.com", "The URL that you want to build a sitemap for")
	flag.Parse()

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	urlLinks, err := link.Parse(resp.Body)

	fmt.Println(urlLinks)
}
