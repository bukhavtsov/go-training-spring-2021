package main

import (
	"fmt"
	"regexp"
)

func isUserName(userName string) bool {
	matched, _ := regexp.MatchString(`^[a-z0-9_]{4,16}$`, userName)
	return matched
}

func main() {
	fmt.Println(isUserName("dima"))
}
