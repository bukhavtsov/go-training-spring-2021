package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// Example 1
	f, err := os.Open("lesson_3/defer/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)    // read number of bites
		os.Stdout.Write(data[:count]) // convert bites to string and print to console
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}

	/*
		// Example 2
		for i := 0; i < 10; i++ {
			defer fmt.Println(i) from 9 to 10
		}
	*/

	/*
		// Example 3
		i := 10
		defer fmt.Println(i) // 20
		i = 20
	*/
}
