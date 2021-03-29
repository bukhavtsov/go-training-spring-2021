package main

import "fmt"

/*
	Description: The first century spans from the year 1 up to and including the year 100,
	The second - from the year 101 up to and including the year 200, etc.

	Given a year, return the century it is in.

	Examples:"
	centuryFromYear(1705)  returns (18)
	centuryFromYear(1900)  returns (19)
	centuryFromYear(1601)  returns (17)
	centuryFromYear(2000)  returns (20)
*/

func century(year int) int {
	if year%100 == 0 {
		return year / 100
	} else {
		return year/100 + 1
	}
}

func main() {
	fmt.Println(century(1705))
	fmt.Println(century(1900))
	fmt.Println(century(1601))
	fmt.Println(century(2000))

}
