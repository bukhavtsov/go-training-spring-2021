package main

import "fmt"

func main() {
	var m = map[string]int{ // map declaration,
		"string1": 1,
		"string2": 2,
		"string3": 3,
	}

	fmt.Println(m)
	fmt.Println(m["string1"]) // get map value
	m["string1"] = 10
	fmt.Println(m["string1"])
	delete(m, "string1")      // delete from map
	fmt.Println(m["string1"]) // 0

	v, ok := m["string2"] // checks value exists or not
	if ok {
		fmt.Println("value exists:", v)
	}

	var m2 map[string]int
	if m2 == nil { // true
		fmt.Println("nil")
	}

	m3 := make(map[string]int, 1)
	m3["example"] = 10  // add to map
	m3["example1"] = 22 // add to map
	fmt.Println(m3["example"])
	fmt.Println(m3["example1"])

	// Maps are not comparable.
	// You can check if they are equal to nil, but you cannot check if two maps have identical keys and values using == or differ using !=.

	// how to iterate by map
	for k,v := range m {
		fmt.Println(k, v)
	}
}
