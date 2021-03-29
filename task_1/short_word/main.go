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
	var length int
	words := strings.Fields(s)
	for _, word := range words {
		if length > len(word) || length == 0 {
			length = len(word)
		}
	}
	return length
}

func main() {
	str := "I am having breakfast now"
	fmt.Println(findShort(str))
}
