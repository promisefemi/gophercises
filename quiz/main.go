package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	csv := flag.String("csv", "problems.csv", "A csv file to import")
	flag.Parse()

	file, err := os.Open(*csv)
	if err != nil {
		exit(fmt.Sprintf("Unable to open file %s", *csv))
	}

	body, err := io.ReadAll(file)
	if err != nil {
		exit("Unfortunately something went wrong")
	}

	fmt.Printf("%s", body)
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
