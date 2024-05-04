package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	myString := "Hello World!"

	// looping through all letters - 0(n)
	newString1, err := reverse(myString)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Println(newString1)

	// using tow pointer method - 0(n/2)
	newString2, err := reverseUsingTwoPointers(myString)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Println(newString2)
}

// reverse taking in a string and reverses the order of the string and returns it
// O(n) - "n" being the string input depending on the length of the string
func reverse(myString string) (string, error) {
	// validate empty string(s)
	if len(strings.ReplaceAll(myString, " ", "")) < 2 {
		return "", errors.New("string is empty")
	}

	// turn string into a slice
	s := strings.Split(myString, "")
	reverseString := ""

	sliceLen := len(s)
	lastKnowIndex := sliceLen - 1

	// reverse loop slice and concatenate characters to make reverse string
	for i := sliceLen; i != 0; i-- {
		// give me the string last know position
		reverseString = reverseString + s[lastKnowIndex]
		lastKnowIndex--
	}

	return reverseString, nil
}

// reverseUsingTwoPointers taking in a string and reverses the order of the string and returns it
// O(n/2) - "n/2" this is due to having two pointers which divides and conquer the string in half
func reverseUsingTwoPointers(myString string) (string, error) {
	// validate empty string(s)
	if len(strings.ReplaceAll(myString, " ", "")) < 2 {
		return "", errors.New("string is empty")
	}

	s := strings.Split(myString, "")
	myReverseStringSlice := make([]string, len(s))

	leftPointer := 0
	rightPointer := len(s) - 1

	for i := 0; i < len(s); i++ {

		// need to find a break point
		if rightPointer < leftPointer {
			break
		}

		myReverseStringSlice[leftPointer] = s[rightPointer]
		myReverseStringSlice[rightPointer] = s[leftPointer]

		// adjust pointer
		leftPointer++
		rightPointer--
	}

	return strings.Join(myReverseStringSlice, ""), nil
}
