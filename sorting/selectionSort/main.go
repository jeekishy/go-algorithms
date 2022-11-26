package main

import "fmt"

// selection sort
// keep track of the lowest number while traversing the array
// then add the lowest number top of array
// O(n^2) - time complexity
// 0(1) - space complexity
func main() {
	toSort := []int{6, 5, 3, 1, 8, 7, 2, 4, 0}
	fmt.Println(toSort)
	fmt.Println(selectionSort(toSort))
}

func selectionSort(toSort []int) []int {
	arrayLength := len(toSort)

	for i := 0; i < arrayLength; i++ {
		// track smallest value
		smallestValueIndex := i

		// loop through array and find the smallest value
		for j := i + 1; j < arrayLength; j++ {
			// check next item is the smallest
			if toSort[smallestValueIndex] > toSort[j] {
				smallestValueIndex = j
			}
		}

		// update the smallest value to top of array
		// swap top smallest index with current
		tempCopy := toSort[i]
		toSort[i] = toSort[smallestValueIndex]
		toSort[smallestValueIndex] = tempCopy
	}

	return toSort
}
