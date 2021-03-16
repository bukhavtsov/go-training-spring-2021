// Example of Pointers
package main

import "fmt"

func main() {
	// & - get the pointer
	// * - dereference

	n := 10
	fmt.Println(&n)    // print pointer of n
	nPtr := &n         // save pointer of n
	fmt.Println(*nPtr) // dereference and print value
}
