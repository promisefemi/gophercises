package main

import (
	"flag"
	"fmt"
	"html/template"
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
	templ := template.Must(template.New("").Parse(cyoa.ParseTemplateHTML("cmd/cyoa/cyoastoryhtml.html")))
	h := cyoa.NewHandler(story,
		cyoa.WithPathFunc(pathFn),
		cyoa.WithTemplate(templ),
	)

	fmt.Printf(" App runninng on http://localhost:%d\n", *port)

	mux := http.NewServeMux()
	mux.Handle("/story/", h)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func pathFn(r *http.Request) string {
	path := r.URL.Path
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}
