package main

import "fmt"

// insertion sort
// sorts each element in an array while traversing the array/list
// it compare with already sorted values and insert new value where required
func main() {
	toSort := []int{6, 5, 3, 1, 8, 7, 2, 4, 0}
	fmt.Println(toSort)
	fmt.Println(insertionSort(toSort))
}

func insertionSort(toSort []int) []int {
	arrayLength := len(toSort)

	// loop through array
	for i := 0; i < arrayLength; i++ {
		// check value is lower than first element
		// move/prepend to front
		if toSort[i] < toSort[0] {
			// make a temp slice
			newTempSlice := make([]int, 0)

			// add new element to start of temp slice
			newTempSlice = append(newTempSlice, toSort[i:i+1]...)

			// add the rest of the slice minus the added element
			newTempSlice = append(newTempSlice, append(toSort[:i], toSort[i+1:]...)...)

			// update slice
			toSort = newTempSlice
			continue
		}

		// when number is not smaller than the first element on the array/slice
		// find where we need to add it
		for j := 1; j < i; j++ {
			if toSort[i] > toSort[j-1] && toSort[i] < toSort[j] {
				// make a temp slice
				newTempSlice := make([]int, 0)
				// keep a temp copy of element to be moved
				elementToMove := toSort[j]

				// copy first chuck of array
				newTempSlice = append(newTempSlice, append(toSort[:j], toSort[i:i+1]...)...)
				// add element to move
				newTempSlice = append(newTempSlice, elementToMove)

				// remove moved element
				toSort = append(toSort[:i], toSort[i+1:]...)

				// add rest of element
				newTempSlice = append(newTempSlice, toSort[len(newTempSlice)-1:]...)

				// update slice
				toSort = newTempSlice
			}
		}
	}

	return toSort
}
