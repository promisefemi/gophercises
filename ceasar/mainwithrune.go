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
	if ch >= 'A' && ch <= 'Z' {
		return rotateWithBase(ch, 'A', delta)
	} else if ch >= 'a' && ch <= 'z' {
		return rotateWithBase(ch, 'a', delta)
	} else {
		return ch
	}
}

func rotateWithBase(ch rune, base, delta int) rune {
	temp := int(ch) - base
	temp = (temp + delta) % 26
	return rune(temp + base)
}
