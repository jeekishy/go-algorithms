package main

import "fmt"

func main() {
	a := []int{6, 7, 0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(findValueAtIndex(a, 0))  // 2
	fmt.Println(findValueAtIndex(a, 6))  // 0
	fmt.Println(findValueAtIndex(a, 1))  // 3
	fmt.Println(findValueAtIndex(a, 5))  // 7
	fmt.Println(findValueAtIndex(a, 2))  // 4
	fmt.Println(findValueAtIndex(a, 78)) // -1*/

	b := []int{7, 0, 1, 2, 3, 4, 5, 6}
	fmt.Println(b)
	fmt.Println(findValueAtIndex(b, 0))  // 1
	fmt.Println(findValueAtIndex(b, 6))  // 7
	fmt.Println(findValueAtIndex(b, 1))  // 2
	fmt.Println(findValueAtIndex(b, 5))  // 6
	fmt.Println(findValueAtIndex(b, 2))  // 3
	fmt.Println(findValueAtIndex(b, 78)) // -1

	c := []int{7, 8, 9, 10, 12, 13, 0, 1, 2, 3, 4, 5, 6}
	fmt.Println(c)
	fmt.Println(findValueAtIndex(c, 0))  // 6
	fmt.Println(findValueAtIndex(c, 6))  // 12
	fmt.Println(findValueAtIndex(c, 1))  // 7
	fmt.Println(findValueAtIndex(c, 5))  // 11
	fmt.Println(findValueAtIndex(c, 2))  // 8
	fmt.Println(findValueAtIndex(c, 78)) // -1

	d := []int{1, 3}
	fmt.Println(d)
	fmt.Println(findValueAtIndex(d, 0)) // -1
	fmt.Println(findValueAtIndex(d, 2)) // -1
	fmt.Println(findValueAtIndex(d, 4)) // -1

	f := []int{3, 4, 5, 6, 7, 8, 1, 2}
	fmt.Println(f)
	fmt.Println(findValueAtIndex(f, 3))  // 0
	fmt.Println(findValueAtIndex(f, 6))  // 3
	fmt.Println(findValueAtIndex(f, 2))  // 7
	fmt.Println(findValueAtIndex(f, 16)) // 7
}

func findValueAtIndex(a []int, target int) int {
	arrayLength := len(a)

	if arrayLength == 1 && target != a[0] {
		return -1
	}

	return binarySearch(a, target, 0, arrayLength-1)
}

// t = 0
// {5,7,0,1,2,3,4,5}
//  0,1,2,3,4,5,6,7

// l: 0, m: 3, h: 7
// {5,7,0,1,2,3,4,5}
//
//			  ^
//	 0,1,2,3,4,5,6,7
//
// check midway = target
// check left side is sorted
// check t is in range
// else
// check eight
// else // left side is not sorted
// check if value falls on the right side
// else
// check value falls on the left side
func binarySearch(a []int, t, lPoint int, hPoint int) int {
	// check if index out of range
	if lPoint < 0 || hPoint > len(a)-1 {
		return -1
	}

	// computer midpoint
	mPoint := (lPoint + hPoint) / 2

	// last element
	if hPoint-lPoint == 0 && t != a[mPoint] {
		return -1
	}

	// check midway has value
	if t == a[mPoint] {
		return mPoint
	}

	if lPoint == mPoint {
		return binarySearch(a, t, mPoint+1, hPoint)
	}

	// check left side is sort
	if a[lPoint] <= a[mPoint] {
		// check value is within range of left side
		if t >= a[lPoint] && t <= a[mPoint-1] {
			return binarySearch(a, t, lPoint, mPoint-1)
		} else {
			return binarySearch(a, t, mPoint+1, hPoint)
		}
	} else { // left side is not sorted
		// check value fall on right side, if so search right side
		if t >= a[mPoint+1] && t <= a[hPoint] {
			return binarySearch(a, t, mPoint+1, hPoint)
		} else {
			// search left side
			return binarySearch(a, t, lPoint, mPoint-1)
		}
	}
}
