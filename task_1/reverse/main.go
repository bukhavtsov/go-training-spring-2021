package main

import "fmt"

/**
 * Description: Complete the solution so that it reverses the string passed into it.
 * Example: "world"  =>  "dlrow"
 */

func reverse(word string) string {
	var result string
	for _, v := range word {
		result = string(v) + result
	}
	return result
}

func main() {
	fmt.Println(reverse("world"))
}
