package main

import (
	"fmt"
	"strings"
)

/*
  Description: An anagram is the result of rearranging the letters of a word to produce a new word (see wikipedia https://en.wikipedia.org/wiki/Anagram).
  Note: anagrams are case insensitive

  Complete the function to return true if the two arguments given are anagrams of each other; return false otherwise.

  Examples:
  "foefet" is an anagram of "toffee"
  "Buckethead" is an anagram of "DeathCubeK"
*/

func isAnagram(test, original string) bool {
	if len(test) != len(original) {
		return false
	}
	for _, testSymbol := range strings.ToLower(test) {
		check := false
		for _, ordinalSymbol := range strings.ToLower(original) {
			if testSymbol == ordinalSymbol {
				check = true
			}
		}
		if !check {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("foefet", "toffee"))
	fmt.Println(isAnagram("Buckethead", "DeathCubeK"))
}
