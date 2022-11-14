package main

import (
	"errors"
	"fmt"
)

// find the first recurring number in an array/slice
// given []int{2,3,4,5,7,8,4,6,1}
// it should return 4

// given []int{1,3,4,5,1,8,4,6,1}
// it should return 1

// given []int{1,3,4,5,8}
// it should return error

func main() {
	a1 := []int{2, 3, 4, 5, 7, 8, 4, 6, 1}
	a1Result, err := findFirstRecurringNumber(a1) // O(n)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("array 1 recurring number: %d\n", a1Result)
	}

	a2 := []int{1, 3, 4, 5, 1, 8, 4, 6, 1}
	a2Result, err := findFirstRecurringNumber(a2) // O(n)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("array 2 recurring number: %d\n", a2Result)
	}

	a3 := []int{1, 3, 4, 5, 8}
	a3Result, err := findFirstRecurringNumber(a3) // O(n)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("array 3 recurring number: %d\n", a3Result)
	}
}

// findFirstRecurringNumber finds the first recurring number in an array
// Time complexity - O(n)
// Memory complexity - O(n)
func findFirstRecurringNumber(a []int) (int, error) {
	loopedResult := map[int]struct{}{} // low memory footprint can also use bitset here instead of struct use bool

	for _, value := range a { // loop is O(n)
		// check value is already in hashmap
		if _, ok := loopedResult[value]; ok { // hashMap lookup is O(1)
			return value, nil
		}

		// else keep value in hashtable
		loopedResult[value] = struct{}{}
	}

	return 0, errors.New("no recurring number found")
}
