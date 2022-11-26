package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"urlshort"
)

func main() {

	yamlFile := flag.String("yml", "", "--yml: src of yaml file")
	jsonFile := flag.String("json", "", "--json: src of json file")
	flag.Parse()
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	handler := urlshort.MapHandler(pathsToUrls, mux)

	// Check both flag: if there is a src for either

	if *yamlFile != "" {
		yamlDataByte, err := readFile(*yamlFile)
		if err != nil {
			panic(err)
		}

		// fmt.Printf("%s", yamlDataByte)

		handler, err = urlshort.YAMLHandler(yamlDataByte, handler)
		if err != nil {
			panic(err)
		}
	} else if *jsonFile != "" {
		jsonDataByte, err := readFile(*jsonFile)
		if err != nil {
			panic(err)
		}
		handler, err = urlshort.YAMLHandler(jsonDataByte, handler)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("You did not add any other file type")
	}

	fmt.Println("Starting the server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		fmt.Println(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func readFile(fileName string) ([]byte, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}
