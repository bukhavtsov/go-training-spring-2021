// Examples of slices
package main

import "fmt"

func changeSlice(slice [5]int) {
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
}

func main() {
	fmt.Println("welcome to lesson 2")
	/*
		// Difference between array and slice
		var x = []int{1, 2, 3}
		fmt.Println(x)
		fmt.Println(len(x)) //3
		fmt.Println(cap(x)) //3

		x = append(x, 4)
		fmt.Println(len(x)) //4
		fmt.Println(cap(x)) //6
	*/

	/*
		// usage of append function
		var x = []int{1, 2, 3}
		var y = []int{4, 5, 6}
		x = append(x, y...)
		x = append(x, 7, 8, 9, 10)
		fmt.Println(x) // [1 2 3 4 5 6, 7, 8, 9, 10]
	*/

	/*
		var x []int
		fmt.Println(x==nil) // true
	*/

	/*
		// make function
		s := make([]int, 0, 10) // make(array, len, cap)
		fmt.Println(s)
		fmt.Println(cap(s)) // 0
		fmt.Println(len(s)) // 10
		var s1 []int
		fmt.Println(cap(s1)) // 0


		// issue with nil
		if s1 == nil {
			fmt.Println("Slice equals nil, but his value is:", s1)
		}

		// reflect.DeepEqual() for comparing slices/arrays


		// Find the max value of the array
		var slice = []int{1, 2, 3, 4, 5}
		max := slice[0]
		for _, item := range slice {
			if item > max {
				max = item
			}
		}
	*/

	/*
		// one more difference between array and slice
		var b = [5]int{5, 6, 7, 4, 5}
		fmt.Println(b)
		changeSlice(b)
		fmt.Println(b)
	*/

	/*
		// slicing slice
			var slice = []int{1, 2, 3, 4, 5, 6, 7}
			y := slice[1:4] // [)
			fmt.Println(y)
			fmt.Println("y cap():", cap(y)) // 6
			fmt.Println("y len():", len(y)) // 3
	*/

	/*
		// slicing slice and changes
		var slice = []int{1, 2, 3, 4, 5, 6, 7}
		y := slice[1:4] // [)
		fmt.Println(y)
		fmt.Println("y cap():", cap(y)) // 6
		fmt.Println("y len():", len(y)) // 3
		fmt.Println("slice:", slice)
		fmt.Println("y:", y)

		slice[2] = 5 // changes in one slice make changes in another
		y[0] = 5
		fmt.Println("slice:", slice)
		fmt.Println("y:", y)
	*/

	/*
		// slicing slice how to lose link
		var slice = []int{1, 2, 3, 4, 5}
		y := slice[1:4] // [)
		fmt.Println(y)
		fmt.Println("y cap():", cap(y)) // 4
		fmt.Println("y len():", len(y)) // 3
		fmt.Println("slice:", slice)
		fmt.Println("y:", y)

		slice = append(slice, 10)
		fmt.Println(slice)


		slice[2] = 5 // now changes in one slice don't make changes in another, because we lose the link
		y[0] = 5
		fmt.Println("slice:", slice) // we can see different slices
		fmt.Println("y:", y)
	*/

	/*
		// How to fix issue with capacity of sliced slice
		x := make([]int, 0, 5)
		x = append(x, 1, 2, 3, 4)
		y := x[:2:2]
		z := x[2:4:4]
		fmt.Println(cap(x), cap(y), cap(z))
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)

		y = append(y, 30, 40, 50)
		x = append(x, 60)
		z = append(z, 70)
		fmt.Println("After appending")
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
	*/

}
