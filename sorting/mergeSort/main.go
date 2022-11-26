package main

import "fmt"

// merge sort
// divide and conquer by dividing the list into half and repeat until we have one item
// take first and second item and compare, then merge and repeat.
// O(n log n) complexity which is linear
func main() {
	toSort := []int{6, 5, 3, 1, 8, 7, 2, 4, 0}
	fmt.Println(toSort)
	fmt.Println(mergeSort(toSort))
}

func mergeSort(toSort []int) []int {
	// base case for recursion
	if len(toSort) == 1 {
		return toSort
	}

	// split array into right and left
	// get halfway
	halfIndex := len(toSort) / 2
	left := toSort[:halfIndex]
	right := toSort[halfIndex:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	// keep track of each side index
	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
			continue
		}

		result = append(result, right[rightIndex])
		rightIndex++
	}

	for ; leftIndex < len(left); leftIndex++ {
		result = append(result, left[leftIndex])
	}
	for ; rightIndex < len(right); rightIndex++ {
		result = append(result, right[rightIndex])
	}

	return result
}
