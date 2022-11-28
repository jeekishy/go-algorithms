package main

import (
	"fmt"
)

func main() {
	a := []int{23, 13, 5, 90, 22, 35, 0, 2}
	fmt.Println(findSecondLargestNumber(a))
}

func findSecondLargestNumber(a []int) int {
	// sort array
	sortedArray := sortArrayDesc(a)
	return sortedArray[1]
}

// -- 5, 2, 1, 8, 0
// is 5 > 2 - Y
// is 2 > 1 - y
// is 1 > 8 - N - swap
// is 1 > 0 Y
// -- 5, 2, 8, 1, 0
// is 5 > 2 - Y
// is 2 > 8 - N - swap
// is 2 > 1 - Y
// is 1 > 0 - y
// -- 5, 8, 2, 1, 0
// is 5 > 8 - N - swap
func sortArrayDesc(a []int) []int {
	arrayLength := len(a)

	for i := 0; i < arrayLength; i++ {
		for j, n := range a {
			// check array length reach
			if (j + 1) == arrayLength {
				break
			}

			// check left & right values
			if n < a[j+1] {
				// swap
				a[j] = a[j+1]
				a[j+1] = n
			}
		}
	}

	return a
}
