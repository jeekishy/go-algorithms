package main

import (
	"fmt"
	"strings"
)

// reverse a string using recursion
func main() {
	fmt.Println(reverseString("ABCD"))
	fmt.Println(reverseString("I AM KING"))
}

func reverseString(str string) string {
	// keep string length
	strLength := len(str)

	// base case
	if strLength < 2 {
		return str
	}

	// split string
	ss := strings.Split(str, "")

	// recursive case
	return reverseString(strings.Join(ss[1:], "")) + strings.Join(ss[:1], "")
}
