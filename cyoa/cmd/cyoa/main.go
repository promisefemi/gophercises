package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/promisefemi/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "The port to start the CYOA web application on")
	fileName := flag.String("file", "gopher.json", "The JSON file with the Choose your own adventure story")
	flag.Parse()
	fmt.Printf("Using the story %s \n", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", &story)
	h := cyoa.NewHandler(story, nil)

	fmt.Printf(" App runninng on http://localhost:%d\n", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
