package main

import "fmt"

/*
 Description: A palindrome is a word, phrase, number, or other
 sequence of characters which reads the same backward or forward.
  This includes capital letters, punctuation, and word dividers.

 Implement a function that checks if something is a palindrome.

 Examples:
 isPalindrome("anna")   ==> true
 isPalindrome("walter") ==> false
 isPalindrome("12321")    ==> true
 isPalindrome("123456")   ==> false
*/

func isPalindrome(str string) bool {
	var reverse string
	for _, char := range str {
		reverse = string(char) + reverse
	}
	return reverse == str
}

func main() {
	fmt.Println(isPalindrome("anna"))
	fmt.Println(isPalindrome("walter"))
	fmt.Println(isPalindrome("12321"))
	fmt.Println(isPalindrome("123456"))

}
