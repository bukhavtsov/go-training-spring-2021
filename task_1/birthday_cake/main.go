package main

func cake(x int, y string) string {
	if x > 1000 {
		return "-1"
	}
	result := 0
	r := []rune(y)
	for i := 0; i < len(r); i++{
		if i % 2 == 0 {
			result += int(r[i])
		} else {
			result += int(r[i]) - 96
		}
	}
	coeff := 0.7
	if float64(x) < float64(x) * coeff {
		return "Fire!!!"
	} else {
		return "That was close!"
	}
}

func main() {
	cake(600, "adc")
}
