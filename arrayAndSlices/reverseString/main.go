package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	myString := "Hello World!"

	newString, err := reverse(myString)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	fmt.Println(newString)
}

// reverse taking in a string and reverses the order of the string and returns it
// O(n) - "n" being the string input depending on the length of the string
func reverse(s string) (string, error) {
	// validate empty string(s)
	if len(strings.ReplaceAll(s, " ", "")) < 2 {
		return "", errors.New("string is empty")
	}

	var sliceOfString []string
	myString := ""

	// turn string into a slice
	sliceOfString = strings.Split(s, "")

	sliceLen := len(sliceOfString)
	lastKnowIndex := sliceLen - 1

	// reverse loop slice and concatenate characters to make reverse string
	for i := sliceLen; i != 0; i-- {
		// give me the string last know position
		myString = myString + sliceOfString[lastKnowIndex]
		lastKnowIndex--
	}

	return myString, nil
}
