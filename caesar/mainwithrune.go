package main

import (
	"fmt"
)

func main() {
	delta := 4
	input := "Promise-femi"

	var ret []rune

	for _, ch := range input {
		ret = append(ret, cipher(ch, delta))
	}
	fmt.Println(string(ret))
}

func cipher(ch rune, delta int) rune {
	// checking if character fall between the rune of capital lettes
	if ch >= 'A' && ch <= 'Z' {
		return rotateWithBase(ch, 'A', delta)
	} else if ch >= 'a' && ch <= 'z' {
		return rotateWithBase(ch, 'a', delta)
	} else {
		return ch
	}
}

func rotateWithBase(ch rune, base, delta int) rune {
	// Resolving character back to the smallest of its rune value
	temp := int(ch) - base
	// calculation position with modulus just like the slice in main.go
	temp = (temp + delta) % 26
	// returnning back to proper rune value of the character
	return rune(temp + base)
}
