package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{0, 3, 4, 51}
	slice2 := []int{4, 6, 50}

	result1 := mergeSortedArrays(slice1, slice2)
	fmt.Println(result1)
}

// merge sorted arrays
func mergeSortedArrays(a1, a2 []int) []int {
	mergedArray := make([]int, 0)

	// merge both array into one
	mergedArray = append(mergedArray, a1[:]...)
	mergedArray = append(mergedArray, a2[:]...)

	// sort merged array
	sort.Sort(sort.IntSlice(mergedArray))

	return mergedArray
}
