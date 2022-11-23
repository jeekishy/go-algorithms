package main

import (
	"fmt"
	"strings"
)

// find if a word or string is Palindromic
// from wikipedia - Some examples of palindromic words are redivider, deified, civic, radar, level, rotor, kayak,
// reviver, racecar, madam, and refer. There are also word-unit palindromes in which the unit of reversal is the
//word ("Is it crazy how saying sentences backwards creates backwards sentences saying how crazy it is?").

func main() {
	if isPalindromic("level") {
		fmt.Println("level is palindromic")
	} else {
		fmt.Println("level is NOT palindromic")
	}

	if isPalindromic("colour") {
		fmt.Println("colour is palindromic")
	} else {
		fmt.Println("colour is NOT palindromic")
	}
}

// check word is palindromic
func isPalindromic(word string) bool {
	// base case
	if len(word) < 2 {
		return true
	}

	// get reverse string
	reverseString := reverseWord(word)

	// compare original word with reverse word
	if word == reverseString {
		return true
	}

	return false
}

// reverse the word using recursion
func reverseWord(word string) string {
	// track word length
	strLength := len(word)

	// base case
	if strLength < 2 {
		return word
	}

	// reverse the word and keep a copy
	breakWord := strings.Split(word, "")

	// next set of words to be handled
	// ABC
	// fun("BC") + A
	//	 |
	// fun(C) + B
	//   |
	//   C
	//  C + B + A
	return reverseWord(strings.Join(breakWord[1:], "")) + breakWord[0]
}
