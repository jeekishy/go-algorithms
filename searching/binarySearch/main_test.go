package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	value    int
	expected int
}

func TestFindIndexOfTarget(t *testing.T) {
	// have a dummy array
	testArray := []int{2, 4, 5, 6, 7, 9, 14}

	// set a range of cases
	testCases := []testCase{
		testCase{value: 7, expected: 4},
		testCase{value: 16, expected: -1},
		testCase{value: 2, expected: 0},
		testCase{value: 14, expected: 6},
		testCase{value: 1, expected: -1},
	}

	// run through each cases and test result
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, findIndexOfTarget(testArray, tc.value))
	}
}
