package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

// 1. Function declaration
func printName1(p Person) {
	fmt.Println(p.name)

}

// 2. Function declaration with a pointer
// & - give address
// * - search value by given address
func printName2(p *Person) {
	fmt.Println(p.name)
}

// 3. Function with multiple params of one type
func multiply(x1, x2 float64) float64 {
	return x1 * x2
}

// 4. Function with variadic Input Parameters and Slices
func getStringOfNumbers(numbers ...int) string {
	var str string
	for _, item := range numbers {
		str = fmt.Sprintf("%s %d", str, item)
	}
	return str
}

// 4. return multiple values
func getStringOfNumbers1(numbers ...int) (string, error) {
	if numbers == nil {
		return "", fmt.Errorf("error: array numbers is empty: %v", numbers)
	}
	var str string
	for _, item := range numbers {
		str = fmt.Sprintf("%s %d", str, item)
	}
	return str, nil
}

// 5. return multiple values and Blank Returns
func getStringOfNumbers3(numbers ...int) (str string, err error) {
	if numbers == nil {
		err = fmt.Errorf("error: array numbers is empty: %v", numbers)
		return // blank Returns — never use these
	}
	for _, item := range numbers {
		str = fmt.Sprintf("%s %d", str, item)
	}
	return
}

// 6. function are values
func add(x1, x2 int) int { return x1 + x2 }

func sub(x1, x2 int) int { return x1 - x2 }

/*

var opMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
	}

	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"two", "+", "three"},
		[]string{"5"},
	}
	// calculate expressions
	for _, expression := range expressions {
		if len(expression) != 3 { // if len is not correct print the error
			fmt.Println("invalid expression:", expression)
			continue
		}
		n1, err := strconv.Atoi(expression[0]) // check that string is int number
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op] // check that operator exists in the map
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		n2, err := strconv.Atoi(expression[2]) // check that string is int number
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(n1, n2) // calculate by existing functions
		fmt.Println(result)      // prints calculation results
	}
*/

type opFuncType func(int, int) int

func main() {
	var opMap = map[string]opFuncType{
		"+": add,
		"-": sub,
	}

	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"two", "+", "three"},
		[]string{"5"},
	}
	// calculate expressions
	for _, expression := range expressions {
		if len(expression) != 3 { // if len is not correct print the error
			fmt.Println("invalid expression:", expression)
			continue
		}
		n1, err := strconv.Atoi(expression[0]) // check that string is int number
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op] // check that operator exists in the map
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		n2, err := strconv.Atoi(expression[2]) // check that string is int number
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(n1, n2) // calculate by existing functions
		fmt.Println(result)      // prints calculation results
	}

}
