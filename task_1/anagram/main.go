package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(test, original string) bool {
	if len(test) != len(original) || test == original{
		return false
	}
	testArray := strings.Split(strings.ToLower(test), "")
	originalArray := strings.Split(strings.ToLower(original), "")
	sort.Strings(testArray)
	sort.Strings(originalArray)
	return fmt.Sprint(testArray) == fmt.Sprint(originalArray)
}

func main(){
	isAnagram("Imad", "Dima")
}
