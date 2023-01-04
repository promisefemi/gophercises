package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	answer := 1
	for _, ch := range input {
		if strings.ToUpper(string(ch)) == string(ch) {
			answer++
		}
	}
	fmt.Println(answer)
}
