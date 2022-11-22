package main

import (
	"fmt"
	"strconv"
	"strings"
)

// given a number, find the sum of its digits
// e.g. n=12 => result=3
// e.g. n=567 => result=18
func main() {

	fmt.Println("sum of digit", 12, "is", sumOfDigitsRecursive(12))
	fmt.Println("sum of digit", 123, "is", sumOfDigitsRecursive(123))
	fmt.Println("sum of digit", 1002, "is", sumOfDigitsRecursive(1002))
	fmt.Println("sum of digit", 1012, "is", sumOfDigitsRecursive(1012))

	fmt.Println("--------------")

	fmt.Println("sum of digit", 12, "is", sumOfDigitsIterative(12))
	fmt.Println("sum of digit", 123, "is", sumOfDigitsIterative(123))
	fmt.Println("sum of digit", 1002, "is", sumOfDigitsIterative(1002))
	fmt.Println("sum of digit", 1012, "is", sumOfDigitsIterative(1012))
}

func sumOfDigitsRecursive(number int) int {
	// keep a string of the number
	s := fmt.Sprint(number)

	// base case for breaking recursion
	if len(s) < 2 {
		return number
	}

	// split strings into digits
	ss := strings.Split(s, "")

	// get first digit from number
	firstDigit, _ := strconv.Atoi(ss[0])

	// add first two digits
	c, _ := strconv.Atoi(strings.Join(ss[1:], ""))

	return firstDigit + sumOfDigitsRecursive(c)
}

func sumOfDigitsIterative(number int) int {
	// convert digit into string
	digit := fmt.Sprint(number)

	// explode into individual string
	digits := strings.Split(digit, "")

	var sum int

	// loop through digits and add them up
	for _, d := range digits {
		c, _ := strconv.Atoi(d)
		sum = sum + c
	}

	return sum
}
