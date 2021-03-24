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
	var (
		a             = strings.Split(s, " ")
		minLengthWord = len(a[0])
	)
	for _, value := range a {
		if minLengthWord > len(value) {
			minLengthWord = len(value)
		}
	}
	return minLengthWord
}

func main() {
	fmt.Println(findShort("String will never be empty and you do not need"))
}
