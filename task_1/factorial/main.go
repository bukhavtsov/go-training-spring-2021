package main

import "fmt"

/*
	In mathematics, the factorial of a non-negative integer n, denoted by n!,
 	is the product of all positive integers less than or equal to n.

 	For example: 5! = 5 * 4 * 3 * 2 * 1 = 120. By convention the value of 0! is 1.

	Write a function to calculate factorial for a given input.
*/

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(5))
}
