package main

import (
	"strings"
)

func isPalindrome(str string) bool {
	if str == "" {
		return false
	}
	array := strings.Split(str, "")
	for i, lastIndex := 0, len(array) - 1; i < len(array); i, lastIndex = i + 1, lastIndex - 1{
		if array[i] != array[lastIndex]{
			return false
		}
	}
	return true
}

func main() {
	isPalindrome("12321")
}
