package main

import "fmt"

/**
 * Description: Complete the solution so that it reverses the string passed into it.
 * Example: "world"  =>  "dlrow"
 */

func reverse(word string) string {
	var result string
	for _, char := range word {
		result = string(char) + result
	}
	return result
}

func main() {
	word := "world"
	word = reverse(word)
	fmt.Println(word)
}
