package main

import "fmt"

// bubble sort is an elementary sort algorithm
// we bubble up the highest number
// easy algorithm but not efficient
// O(n^2) - time complexity
// 0(1) - space complexity
func main() {
	toSort := []int{6, 5, 3, 1, 8, 7, 2, 4}
	fmt.Println(toSort)
	fmt.Println(bubbleSort(toSort))
}

// [6, 5, 3, 1, 8, 7, 2, 4]
//  0, 1, 2, 3, 4, 5, 6, 7
// expected result [1, 2, 3, 4, 5, 6, 7, 8]
func bubbleSort(toSort []int) []int {
	arrayLength := len(toSort)

	for i := 0; i < arrayLength; i++ {
		for j, n := range toSort {
			// check end of array
			if (j + 1) == arrayLength {
				break
			}

			// sort element
			rightElement := toSort[j+1]
			if n > rightElement {
				toSort[j] = rightElement
				toSort[j+1] = n
				continue
			}
		}
	}

	return toSort
}
