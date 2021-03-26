package main

import (
	"fmt"
	"time"
)

// types
type Developer struct {
	Name         string
	PrimarySkill string
}

// own types
type MyPersonalInt int

// define type string
func (d Developer) String() string {
	return fmt.Sprintf("developer name: %s\ndeveloper PrimarySkill: %s\n------------\n", d.Name, d.PrimarySkill)
}

// Type Declarations Aren’t Inheritance. Types just a documentation
type A int
type B A


type MailCategory int
const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

type BitField int

const (
	Field1 BitField = 1 << iota // assigned 1
	Field2                      // assigned 2
	Field3                      // assigned 4
	Field4                      // assigned 8
)


type Counter struct {
	total             int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) Increment1() {
	c.total++
	c.lastUpdated = time.Now()
}

func main() {
	var a MyPersonalInt
	fmt.Println(a)

	dev := Developer{
		Name:         "Artsiom",
		PrimarySkill: "Golang",
	}
	fmt.Println(dev)

}
