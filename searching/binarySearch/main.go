package main

import "fmt"

// binary search uses the approach of divide and conquer
// we need to find the halfway point at any given time and check
// if the value is in the left or right and repeat until we find the
// target number we are looking for
// O(n Log N) - time complexity
func main() {
	// given a sorted array
	a := []int{2, 7, 10, 13, 30, 56, 89, 90, 100, 120, 130, 145}

	// find the index of the target value
	fmt.Println("sorted array:", a)
	fmt.Printf("index of %d is %d\n", 7, findIndexOfTarget(a, 7))
	fmt.Printf("index of %d is %d\n", 10, findIndexOfTarget(a, 10))
	fmt.Printf("index of %d is %d\n", 13, findIndexOfTarget(a, 13))
	fmt.Printf("index of %d is %d\n", 30, findIndexOfTarget(a, 30))
	fmt.Printf("index of %d is %d\n", 56, findIndexOfTarget(a, 56))
	fmt.Printf("index of %d is %d\n", 89, findIndexOfTarget(a, 89))
	fmt.Printf("index of %d is %d\n", 90, findIndexOfTarget(a, 90))
	fmt.Printf("index of %d is %d\n", 100, findIndexOfTarget(a, 100))
	fmt.Printf("index of %d is %d\n", 120, findIndexOfTarget(a, 120))
	fmt.Printf("index of %d is %d\n", 130, findIndexOfTarget(a, 130))
	fmt.Printf("index of %d is %d\n", 145, findIndexOfTarget(a, 145))
	fmt.Printf("------ SHOULD NOT FIND INDEX --------\n")
	fmt.Printf("index of %d is %d\n", 135, findIndexOfTarget(a, 135))
	fmt.Printf("index of %d is %d\n", 35, findIndexOfTarget(a, 35))
	fmt.Printf("index of %d is %d\n", 52, findIndexOfTarget(a, 52))

}

func findIndexOfTarget(a []int, target int) int {
	// keep the length of array
	arrayLength := len(a)

	// basic check
	if arrayLength == 1 && target != a[0] {
		return -1
	}

	return binarySearch(a, target, 0, arrayLength-1)
}

func binarySearch(a []int, target, lPoint, hPoint int) int {
	// compute midway point
	mPoint := (lPoint + hPoint) / 2

	// check when only 1 value is remaining
	if hPoint-lPoint == 0 && target != a[mPoint] {
		return -1
	}

	// base case
	if target == a[mPoint] {
		return mPoint
	}

	// recursive case
	// search left side when value is less
	if target < a[mPoint] {
		return binarySearch(a, target, lPoint, mPoint-1)
	}

	// search right side if value is more than midpoint
	return binarySearch(a, target, mPoint+1, hPoint)
}
