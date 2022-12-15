package main

import (
	"fmt"
)

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
func main() {
	a := []int{0, 2, 5, 13, 22, 23, 35, 90}
	fmt.Println(twoSum(a, 37))
}

func twoSum(a []int, sumToFind int) []int {
	aLen := len(a)

	// loop through slice
	for i := 0; i < aLen; i++ {
		// assign left and right pointers
		lPointer := i + 1
		rPointer := aLen - 1

		for j := i + 1; j < aLen; j++ {
			// check left pointer
			if a[i]+a[lPointer] == sumToFind {
				return []int{i, lPointer}
			}

			// check right pointer
			if a[i]+a[rPointer] == sumToFind {
				return []int{i, rPointer}
			}

			// check pointer over crossing each other
			if rPointer <= lPointer {
				break // end check
			}

			// else increment pointers
			lPointer++
			rPointer--
		}
	}

	return []int{}
}
