package main

import "fmt"

/**
 * Description: Complete the solution so that it reverses the string passed into it.
 * Example: "world"  =>  "dlrow"
 */

func reverse(word string) string {
	var reverse string
	for _, v := range word {
		reverse = string(v) + reverse
	}
	return reverse
}

func main() {
	fmt.Println(reverse("world"))
}
