package main

import (
	"fmt"
	"strings"
)

func caesar(r rune, shift int) rune {
	// Shift character by specified number of places
	// If beyond range, shift backwards or forwards
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}

func main() {
	value := "test"
	fmt.Println(value)

	// Test the caesar method in a func argument to strings.Map
	value2 := strings.Map(func(r rune) rune {
		return caesar(r, 18)
	}, value)
	value3 := strings.Map(func(r rune) rune {
		return caesar(r, -18)
	}, value2)
	fmt.Println(value2, value3)

	value4 := strings.Map(func(r rune) rune {
		return caesar(r, 1)
	}, value)

	value5 := strings.Map(func(r rune) rune {
		return caesar(r, -1)
	}, value4)

	fmt.Println(value4, value5)
}
