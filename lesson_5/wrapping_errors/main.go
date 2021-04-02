package main

import (
	"fmt"
)

type errorFirst struct{}

func (e errorFirst) Error() string {
	return "Error First happened"
}

func main() {
	e1 := errorFirst{}
	e2 := fmt.Errorf("error 2: %w", e1)
	e3 := fmt.Errorf("error 3: %w", e2)
	fmt.Println(e2)
	fmt.Println(e3)
}
