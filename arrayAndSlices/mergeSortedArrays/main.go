package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{0, 3, 4, 51}
	slice2 := []int{4, 6, 50}

	// using go helper methods
	result1 := mergeSortedArrays(slice1, slice2)
	fmt.Println(result1)

	// using pointers
	result2 := mergeSortedArraysMyTake(slice1, slice2)
	fmt.Println(result2)
}

// merge sorted arrays
func mergeSortedArrays(firstArray, secondArray []int) []int {
	mergedArray := make([]int, 0)

	// merge both array into one
	mergedArray = append(mergedArray, firstArray[:]...)
	mergedArray = append(mergedArray, secondArray[:]...)

	// sort merged array
	sort.Sort(sort.IntSlice(mergedArray))

	return mergedArray
}

func mergeSortedArraysMyTake(firstArray, secondArray []int) []int {
	totalArraySize := len(firstArray) + len(secondArray)
	mergedArray := make([]int, totalArraySize)

	insertIndex := 0
	firstArrayPointer := 0
	secondArrayPointer := 0

	for insertIndex < totalArraySize {

		if firstArrayPointer == len(firstArray) {
			mergedArray[insertIndex] = secondArray[secondArrayPointer]
			insertIndex++
			secondArrayPointer++
			continue
		}

		if secondArrayPointer == len(secondArray) {
			mergedArray[insertIndex] = firstArray[firstArrayPointer]
			insertIndex++
			firstArrayPointer++
			continue
		}

		if firstArray[firstArrayPointer] < secondArray[secondArrayPointer] {
			mergedArray[insertIndex] = firstArray[firstArrayPointer]
			insertIndex++
			firstArrayPointer++
			continue
		}

		mergedArray[insertIndex] = secondArray[secondArrayPointer]
		insertIndex++
		secondArrayPointer++
	}

	return mergedArray
}
