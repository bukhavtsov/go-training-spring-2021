// Example of arrays
package main

import "fmt"

func main() {
	fmt.Println("welcome to lesson 2")
	// Define a simple array
	var x [7]int
	fmt.Println(x) // [0 0 0 0 0 0 0]

	/*
		// Array literal
		var x =  [4]int{1,2,3,4}
		fmt.Println(x) // [1 2 3 4]

	*/

	/*
		// sparse array
		var x = [10]int{1, 2, 3, 7: 4, 9: 6}
		fmt.Println(x) // [1 2 3 0 0 0 0 4 0 6]
	*/

	/*
		// you can leave the number of elements and use ... instead
		var x = [...]int{1,2,3,4,5,6,7,8,9}
		fmt.Println(x) // [1 2 3 4 5 6 7 8 9]
	*/

	/*
		// compare arrays by == and !=
		var x = [3]int{7,7,7}
		var y = [...]int{7,7,7}
		var z = [...]int{77,77,77}
		fmt.Println(x == y) // true
		fmt.Println(x == z) // false
	*/

	/*
		// len - size of array
		// cap - number of allocated elements(equal to len)
			var x = [3]int{7,7,7}
			fmt.Println(len(x)) // 3
			fmt.Println(cap(x)) // 3
	*/

	/* Loop through an array with range statement
	var x = [3]int{7, 7, 7}
	for i, element := range x {
		fmt.Println(i, element)
	}
	*/
}
