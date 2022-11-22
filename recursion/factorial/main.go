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
	// base case - stops the recursion
	if number < 2 {
		return 1
	}

	n := number

	// recursive case - n * (n - 1)
	n = n * getFactorialRecursive(n-1)

	return n
}

func getFactorialIterative(number int) int {
	// keep track of result
	result := number
	lastNumber := number

	// n * (n-1) * .....
	for i := 1; i < number; i++ {
		result = result * (lastNumber - 1)
		lastNumber--
	}

	return result
}
