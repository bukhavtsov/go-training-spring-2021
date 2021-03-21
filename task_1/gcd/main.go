package main

func compute(x, y int) int {
	if y == 0 {
		return x
	} else {
		return compute(y,x%y)
	}
}

func main() {
	compute(9, 3)
}