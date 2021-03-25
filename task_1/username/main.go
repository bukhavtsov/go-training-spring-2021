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
	matched, _ := regexp.MatchString(`^[a-z_\d]{4,16}$`, username)
	return matched
}

func main() {
	fmt.Println(isUsername("____"))
	fmt.Println(isUsername("1111"))
	fmt.Println(isUsername("Eugen"))
}
