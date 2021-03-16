// Example of logical operators
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 30

	// If and else
	randNum := rand.Intn(max-min+1) + min
	if randNum < 20 {
		fmt.Println("randNum less then 20. value is:", randNum)
	} else {
		fmt.Println("randNum higher 20. value is:", randNum)
	}



	// If and else with a short statement
	if randNum = rand.Intn(max-min+1) + min; randNum < 25 {
		fmt.Println("randNum less then 25. value is:", randNum)
	} else {
		fmt.Println("randNum higher 25. value is:", randNum)
	}

}
