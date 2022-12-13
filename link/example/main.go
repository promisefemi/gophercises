package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/promisefemi/gophercises/link"
)

func main() {

	exampleHtml := flag.String("file", "ex1.html", "Html File to read from")
	flag.Parse()

	// strReader := strings.NewReader(exampleHtml1)

	openedExample, err := os.Open(*exampleHtml)

	if err != nil {
		panic(err)
	}

	htmlLinks, err := link.Parse(openedExample)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", htmlLinks)
}
