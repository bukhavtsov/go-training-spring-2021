package main

import (
	"fmt"
	"regexp"
)

/*
	Write a simple regex to validate a username. Allowed characters are:
	lowercase letters,
	numbers,
	underscore
	Length should be between 4 and 16 characters (both included).
*/

func isUsername(username string) bool {
	matched, err := regexp.MatchString("^([0-9a-z]|[0-9a-z]+_?[0-9a-z]+){4,16}$", username)
	if err != nil {
		return false
	}
	return matched
}

func main() {
	name1 := "Alexandr"
	name2 := "alexandr"
	name3 := "alexandr_59"
	name4 := "alexandr59"
	name5 := "Alexandr_59"
	fmt.Println(isUsername(name1))
	fmt.Println(isUsername(name2))
	fmt.Println(isUsername(name3))
	fmt.Println(isUsername(name4))
	fmt.Println(isUsername(name5))
}
