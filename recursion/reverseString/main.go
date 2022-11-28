package main

import (
	"fmt"
	"strings"
)

// reverse a string using recursion
func main() {
	fmt.Println(reverseString("ABCD"))
	fmt.Println(reverseString("I am Driving"))

	fmt.Println(reverseString2("Awesome"))
	fmt.Println(reverseString2("Fifa 2022 is on"))
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

func reverseString2(str string) string {
	strLength := len(str)

	// base case and check
	if strLength == 1 {
		return str
	}

	// split string
	ss := strings.Split(str, "")

	// add last character and run the other recursively
	return ss[strLength-1] + reverseString2(str[:strLength-1])
}
