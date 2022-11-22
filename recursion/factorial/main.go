package main

import "fmt"

// factorial of 4
// n = 4 * 3 * 2 * 1
// 0! = 1
// 1! = 1
// n = n * n-1 ... = n\
func main() {
	// using recursion
	fmt.Println(getFactorialRecursive(4))
	fmt.Println(getFactorialRecursive(5))
	fmt.Println(getFactorialRecursive(8))

	// using a loop
	fmt.Println(getFactorialIterative(4))
	fmt.Println(getFactorialIterative(5))
	fmt.Println(getFactorialIterative(8))
}

// use recursion - uses stacks
// prone to stack overflow if number is big
func getFactorialRecursive(number int) int {
	// check zero // 0 1
	if number < 2 {
		return 1
	}

	return number * getFactorialRecursive(number-1)
}

func getFactorialIterative(number int) int {
	// check zero // 0 1
	if number < 2 {
		return 1
	}

	n := number
	lastNumber := number

	for i := 1; i < number; i++ {
		n = n * (lastNumber - 1)
		lastNumber--
	}

	return n
}