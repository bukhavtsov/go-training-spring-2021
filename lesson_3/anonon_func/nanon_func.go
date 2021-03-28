package main

import "fmt"

// closure example 1
func createPrintNameFunc(name string) func() {
	return func() {
		fmt.Println("name is", name)
	}
}

// closure example 2
func createIncrementer(n int) func(int) int {
	return func(number int) int {
		return n + number
	}
}

// closure example 3
func urlGenerator(domain string) func(string) string {
	return func(url string) string {
		return fmt.Sprintf("https://%s.%s", url, domain)
	}
}

// func as a param check lesson 2 sorting folder

func main() {

	// anonymous func example 1
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(i)
	}

	printName := createPrintNameFunc("Jacob")
	printName()

	incrementNumberToOne := createIncrementer(1) // increment some number to value that added to input
	fmt.Println(incrementNumberToOne(10))        // func incrementNumberToOne closures input param that given into  createIncrementer func
	fmt.Println(incrementNumberToOne(22))

	incrementNumberToFour := createIncrementer(4)
	fmt.Println(incrementNumberToFour(10)) // 14

	comURL := urlGenerator("com")
	byURL := urlGenerator("by")
	fmt.Println(comURL("google"))
	fmt.Println(byURL("google"))
}
