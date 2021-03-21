package main

import (
	"strings"
)

func findShort(s string) int {
	var array []string = strings.Split(s, " ")
	min := len(array[0])
	for _, value := range array {
		if min > len(value){
			min = len(value)
		}
	}
	return min
}

func main() {
	findShort("Dima d asd asdasda")
}
