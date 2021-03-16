// Example of loops
package main

import "fmt"

func main() {
	// For
	for i := 0; i < 5; i++ {
		fmt.Println("for: Message example, i:", i)
	}
	// For is Go's "while"
	i := 0
	for i < 5 {
		fmt.Println("for continued: Message example, i:", i)
		i++
	}

	// Forever
	/*
		for {

		}
	*/
}
