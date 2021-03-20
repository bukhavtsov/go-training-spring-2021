package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// All sorting algorithms in the Go sort package make O(n log(n)) comparisons
func main() {
	// Sort a slice of ints, float64s or strings
	var slice = []int{1, 2, 3, 4, 5}

	// shuffle elements from the array
	rand.Seed(time.Now().UnixNano())

	var swap = func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	}
	rand.Shuffle(len(slice), swap)

	fmt.Println(slice)
	sort.Ints(slice)
	fmt.Println(slice)

	// Sort with custom comparator
	product := []struct {
		name  string
		price int
	}{
		{"phone", 500},
		{"TV", 1500},
		{"car", 20_000},
		{"burger", 2},
		{"bicycle", 200},
	}
	fmt.Println(product)

	sort.SliceStable(product, func(i, j int) bool {
		return product[i].price < product[j].price
	})
	fmt.Println(product) // [{burger 2} {bicycle 200} {phone 500} {TV 1500} {car 20000}]

	// Sorting map keys
	m := map[string]int{
		"phone":   500,
		"tv":      1500,
		"car":     20_000,
		"burger":  2,
		"bicycle": 200,
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}

	// binary search
	var testValues [100]int
	for i := 0; i < 100; i++ {
		testValues[i] = i
	}

	searchNumber := 50

	i := sort.Search(len(testValues), func(i int) bool { return testValues[i] >= searchNumber })
	if i < len(testValues) && testValues[i] == searchNumber {
		fmt.Printf("found %d at index %d in %v\n", searchNumber, i, testValues)
	} else {
		fmt.Printf("%d not found in %v\n", searchNumber, testValues)
	}
}
