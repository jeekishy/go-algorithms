package main

import "fmt"

// factorial of 4
// n = 4 * 3 * 2 * 1
// 0! = 1
// 1! = 1
// n = n * n-1 ... = n\
func main() {
	// using recursion
	fmt.Println(getFactorial(4))
	fmt.Println(getFactorial(5))
	fmt.Println(getFactorial(8))

	// using a loop
	fmt.Println(getFactorialLoop(4))
	fmt.Println(getFactorialLoop(5))
	fmt.Println(getFactorialLoop(8))
}

// use recursion - uses stacks
func getFactorial(number int) int {
	// check zero // 0 1
	if number < 2 {
		return 1
	}

	return number * getFactorial(number-1)
}

func getFactorialLoop(number int) int {
	// check zero // 0 1
	if number < 2 {
		return 1
	}

	n := number
	lastNumber := number

	for i := 1; i < number; i++ {
		// 4 * (4 - 1) - 4 x 3
		// 12 * (3 -1 ) - 12 x 2
		// 24 * (2-1) - 24 * 1
		n = n * (lastNumber - 1)
		lastNumber--
	}

	return n
}
