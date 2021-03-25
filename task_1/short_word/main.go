package main

import (
	"fmt"
	"strings"
)

/*
	given a string of words, return the length of the shortest word(s).
	String will never be empty and you do not need to account for different data types.
*/

func findShort(s string) int {
	words := strings.Split(s, " ")
	minLength := len(words[0])
	var length int
	for _, v := range words {
		length = len(v)
		if length < minLength {
			minLength = length
		}
	}
	return minLength
}

func main() {
	fmt.Println(findShort("given a string of words, return the length of the shortest word"))
}
