package main

import (
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	max := math.Sqrt(float64(n)) + 1
	for i := 2; i < int(max); i++{
		if n % i == 0{
			return false
		}
	}
	return true
}

func main() {
	isPrime(115)
}
