package main

import "fmt"

// find the missing number in a fixed length array which has been sorted
// the array number starts with could start with 1 and have 9 elements and the
// maximum number/or last number in the array would be 10
func main() {
	sampleArray := [9]int{1, 2, 3, 4, 5, 6, 8, 9, 10}
	sampleArray2 := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	sampleArray3 := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(findMissingNumber(sampleArray))
	fmt.Println(findMissingNumber(sampleArray2))
	fmt.Println(findMissingNumber(sampleArray3))
}

// finds the missing number in an array - O(n)
func findMissingNumber(a [9]int) int {
	expectedNumber := 1

	for _, n := range a {
		if expectedNumber != n {
			return expectedNumber
		}

		expectedNumber++
	}

	return 0
}
