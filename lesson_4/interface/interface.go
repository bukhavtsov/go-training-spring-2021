package main

import (
	"math"
)

// Example 2
/*
type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("Processing....")
	}
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	// get data from somewhere
	c.L.Process("data")
}
*/

/*
// Example 6
func doThings(i interface{}) {
	switch j := i.(type) {
	case nil:
		// i is nil, type of j is interface{}
	case int:
		// j is of type int
	case MyInt:
		// j is of type MyInt
	case io.Reader:
		// j is of type io.Reader
	case string:
		// j is a string
	case bool, rune:
		// i is either a bool or rune, so j is of type interface{}
	default:
		// no idea what i is, so j is of type interface{}
	}
}
*/
type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	/*
	// Example 1
	c1 := circle{5}
	r1 := rectangle{7,8}
	shapes := []shape{c1, r1}
	for _,s := range shapes{
		fmt.Println(s)
	}
	*/
	/* Example 2
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
	*/

	/*
		// Example 3
		var s *string
		fmt.Println(s == nil) // prints true
		var i interface{}
		fmt.Println(i == nil) // prints true
		i = s
		fmt.Println(i == nil) // prints false
	*/

	/*
		// Example 4
		var i interface{}
		i = 20
		i = "hello"
		i = struct {
			FirstName string
			LastName string
		} {"Fred", "Fredson"}
	*/

	/*
			// Example 5
			var i interface{}
		    var mine MyInt = 20
		    i = mine
		    i2 := i.(MyInt)
		    fmt.Println(i2 + 1)
	*/

	/*
		i2, ok := i.(int)
		if !ok {
			return fmt.Errorf("unexpected type for %v",i)
		}
		fmt.Println(i2 + 1)
	*/
}
