// Example of structures
package main

type person struct {
	name string
	age  int
	male bool
}

type student struct {
	name string
	age  int
	male bool
}

func main() {
	/*
		// Example_1
		// Anonymous functions
		a := struct {
			name string
			age  int
			male bool
		}{
			"Jacob",
			14,
			true,
		}

		fmt.Println(a)
	*/
	/*
		// Comparable
		// Example_2
		// If structure consist of primitive data types, then you can compare values of this structures
		p1 := person{
			name: "Robert",
			age: 18,
			male: true,
		}

		p2 := person{
			name: "Robert",
			age: 18,
			male: true,
		}


		p3 := person{
			name: "Tom",
			age: 21,
			male: true,
		}

		isEqual := p1 == p2
		isNotEqual := p1 == p3

		fmt.Println(isEqual) // true
		fmt.Println(isNotEqual) // false
	*/

	/*
		// Example_3
		// If structures have different type, but identical fields they are not comparable, but we can convert one type to another
		p := person{
			name: "Martin",
			age:  20,
			male: true,
		}
		s := student{
			name: "Martin",
			age:  20,
			male: true,
		}

		// isNotEqual := p == s // We'll get a panic, because different types
		isEqual := person(s) == p
		fmt.Println(isEqual) // true
	*/

	/*
		// Example_4
		// If struct and anonymous struct have equal fields, they are comparable
		p := person{
			name: "Martin",
			age:  20,
			male: true,
		}
		m := struct {
			name string
			age  int
			male bool
		}{
			name: "Martin",
			age:  20,
			male: true,
		}

		isEqual := p == m
		fmt.Println(isEqual) // true
	*/
}
