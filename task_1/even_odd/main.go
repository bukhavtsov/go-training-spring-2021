package main

import "fmt"

/*
 Description:
 Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.
*/

func evenOrOdd(number int) string {
	result := "Even"
	if number%2 != 0 {
		result = "Odd"
	}
	return result
}

func main() {
	fmt.Println(evenOrOdd(4))
}
