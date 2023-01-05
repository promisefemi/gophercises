package main

import (
	"fmt"
	"strings"
)

func main() {
	delta := 4
	input := "Promise-femi"

	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := strings.ToUpper(alphabetLower)
	ret := ""

	for _, ch := range input {
		switch {
		case strings.IndexRune(alphabetLower, ch) >= 0:
			ret = ret + string(rotate(ch, delta, []rune(alphabetLower)))
		case strings.IndexRune(alphabetUpper, ch) >= 0:
			ret = ret + string(rotate(ch, delta, []rune(alphabetUpper)))
		default:
			ret = ret + string(ch)
		}
	}
	fmt.Println(ret)
}

func rotate(text rune, shiftPosition int, key []rune) rune {
	textPosition := strings.IndexRune(string(key), text)

	if textPosition < 0 {
		panic("A text you are trying to rotate cannot be found in the key you presented")
	}

	// Simple solution by adding and subtracting positions
	// if textPosition+shiftPosition >= len(key) {
	// 	remainder := int(len(key) - (textPosition + shiftPosition))
	// 	return key[remainder]
	// } else {
	// 	return key[textPosition+shiftPosition]
	// }

	// Using loops
	// for i := 0; i <= shiftPosition; i++ {
	// 	textPosition++
	// 	if textPosition >= len(key) {
	// 		textPosition = 0
	// 	}
	// }

	// Using Modulus
	textPosition = (textPosition + shiftPosition) % len(key)
	return key[textPosition]
}
