package main

import (
	"fmt"
	"strings"
)

func reverse(word string) string {
	if len(word) == 0{
		return "-1"
	}
	var array []string = strings.Split(word, "")
	for i, lastIndex := 0, len(array) - 1; i < lastIndex; i, lastIndex = i + 1, lastIndex - 1{
		array[i], array[lastIndex] = array[lastIndex], array[i]
	}
	return fmt.Sprint(array)
}

func main() {
	reverse("Dima")
}
