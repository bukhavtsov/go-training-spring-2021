// Example of basic data types
package main

import (
	"fmt"
)

func main() {
	var (
		number               int
		s                    string
		boolean              bool
		float                = 64.0e5
		complex              complex64
		runeUnicodeCharacter = 'a'
		runeLiteralEightBit  = '\141'
		runeLiteralHEX8      = '\x61'
		runeLiteralHEX16     = '\u0061'
		runeLiteralHEX32     = '\U00000061'
		newLine              = '\n'
		tab                  = '\t'

		message = `
start of message
continue on new line
end new
new 
....`

		simpleMessage = "start of message\ncontinue on new line\nend new\nnew "
	)

	fmt.Println(number)
	fmt.Println(s)
	fmt.Println(boolean)
	fmt.Println(float)
	fmt.Println(complex)
	fmt.Println(runeUnicodeCharacter)
	fmt.Println(runeLiteralEightBit)
	fmt.Println(runeLiteralHEX8)
	fmt.Println(runeLiteralHEX16)
	fmt.Println(runeLiteralHEX32)
	fmt.Println(newLine)
	fmt.Println(tab)
	fmt.Println(message)
	fmt.Println(simpleMessage)
}
