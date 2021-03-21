package main

import (
	"fmt"
	"sort"
)

func isValidArray(numbers []int) bool{
	sort.Ints(numbers)
	for i := 0; i < len(numbers); i++{
		index := sort.SearchInts(numbers, i)
		if index == len(numbers) || i != numbers[index]{
			return false
		}
	}
	return true
}

func getDuplicate(numbers []int) int {
	if !isValidArray(numbers){
		return -1
	}
	var number int = 0
	for i := 0; i < len(numbers); i++{
		for j := 1 ; j < len(numbers); j++{
			if numbers[i] == numbers[j] {
				number = numbers[i]
				return number
			}
		}
	}
	return number
}

func main() {
	array := []int{
		3, 2, 5, 1, 3, 4,
	}
	fmt.Println(getDuplicate(array))
}
