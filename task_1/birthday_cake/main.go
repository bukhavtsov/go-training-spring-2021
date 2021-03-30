package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	It's your Birthday. Your colleagues buy you a cake. The numbers of candles on the cake is provided (x). Please note this is not reality, and your age can be anywhere up to 1,000. Yes, you would look a mess.
	As a surprise, your colleagues have arranged for your friend to hide inside the cake and burst out. They pretend this is for your benefit, but likely it is just because they want to see you fall over covered in cake. Sounds fun!
	When your friend jumps out of the cake, he/she will knock some of the candles to the floor. If the number of candles that fall is higher than 70% of total candles (x), the carpet will catch fire.
	You will work out the number of candles that will fall from the provided string (y). You must add up the character ASCII code of each even indexed (assume a 0 based indexing) character in y, with the alphabetical position of each odd indexed character in y to give the string a total.
	example: abc --> a=97, b=2, c=99 --> y total = 198.
	If the carpet catches fire, return "Fire!", if not, return "That was close!".

	How to convert character to ASCII and back: https://www.socketloop.com/tutorials/golang-how-to-convert-character-to-ascii-and-back
*/
func main() {
	fmt.Println(cake(25, "abc"))
}
func cake(x int, y string) string {
	const alphabet = "abcdefghijklmnopqrstunwxyz"
	var total int

	for i, v := range strings.ToLower(y) {
		if i%2 == 0 {
			number, _ := strconv.Atoi(string(v))
			total = total + number
		} else {
			total = total + strings.IndexAny(alphabet, string(v)) + 1
		}
	}
	return isFire(total, x)
}

func isFire(total, x int) string {
	result := "That was close!"
	if (x*70)/100 < total {
		result = "Fire!"
	}
	return result
}
