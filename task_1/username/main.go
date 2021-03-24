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
	matched, _ := regexp.MatchString(`^[a-z_\d\S]{4,16}$`, username)
	return matched
}

func main() {
	fmt.Println(isUsername("Sla_va_"))
	fmt.Println(isUsername("Sl_ava"))
	fmt.Println(isUsername("Sl ava"))
	fmt.Println(isUsername("Sl123_ava "))
	fmt.Println(isUsername("1123 Sla_va_"))
}
