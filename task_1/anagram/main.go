package main

import (
	"fmt"
	"sort"
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
	return sortString(strings.ToLower(test)) == sortString(strings.ToLower(original))
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	fmt.Println(isAnagram("foefet", "toffee"))
	fmt.Println(isAnagram("Buckethead", "DeathCubeK"))
}
