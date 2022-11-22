package main

import "fmt"

// implement Fibonacci sequence where last two digits totals to the next number
// e.g. 0, 1 = 0 + 1 = 1 => 0, 1, 1 = 1 + 1 = 2 => 0, 1, 1, 2
// e.g. 0, 1, 1, 2, 3, 5, 8, 13
func main() {

	fmt.Println(fibonacciRecursive(0))
	fmt.Println(fibonacciRecursive(1))
	fmt.Println(fibonacciRecursive(2))
	fmt.Println(fibonacciRecursive(3))
	fmt.Println(fibonacciRecursive(4))
	fmt.Println(fibonacciRecursive(5))
	fmt.Println(fibonacciRecursive(6))
	fmt.Println(fibonacciRecursive(7))

	fmt.Println("------------")

	// Iterative solution
	fmt.Println(fibonacciIterative(0))
	fmt.Println(fibonacciIterative(1))
	fmt.Println(fibonacciIterative(2))
	fmt.Println(fibonacciIterative(3))
	fmt.Println(fibonacciIterative(4))
	fmt.Println(fibonacciIterative(5))
	fmt.Println(fibonacciIterative(6))
	fmt.Println(fibonacciIterative(7))
	fmt.Println(fibonacciIterative(8))
	fmt.Println(fibonacciIterative(9))
}

// O(2^N) - exponential time - very bad way of doing it
// return fibonacci number at provided index
// e.g. when index is 4  should return 3 -- 0, 1, 1, 2, 3, 5, 8
// e.g. when index is 2  should return 1 -- 0, 1, 1, 2, 3, 5, 8
func fibonacciRecursive(index int) int {
	// base case - stops the recursion
	// 0 - index will always be 0
	if index < 2 {
		return index
	}

	// add previous fibonacci number with the before previous fibonacci number
	return fibonacciRecursive(index-1) + fibonacciRecursive(index-2)
}

// o(n)
// use a loop to generate fibonacci number
// e.g. 0+1=1 => 1+1=2 => 1+2=3 => 2+3=5
// pattern is second number + last fibonacci = next fibonacci
func fibonacciIterative(index int) int {
	// validate first two index
	if index < 2 {
		return index
	}

	// fibonacci sequence starter
	// since we are handling first two index start with second and third fibonacci numbers
	lastFibonacci := 1
	currentFibonacci := 1

	for i := 2; i < index; i++ {
		// keep a copy of current fibonacci number
		tempCurrent := currentFibonacci
		currentFibonacci = lastFibonacci + currentFibonacci // 5
		// update last fibonacci number
		lastFibonacci = tempCurrent
	}

	return currentFibonacci
}
