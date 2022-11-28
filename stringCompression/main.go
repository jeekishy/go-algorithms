package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compressString([]byte("aabbccczzzz")))
}

func compressString(chars []byte) string {
	if len(chars) == 1 {
		return string(chars[0])
	}

	var result string

	// loop through characters
	j := 0
	for i, v := range chars {
		if v != chars[j] || i == len(chars)-1 {
			// concatenate string result
			result = result + string(chars[j])

			// include repeated string count at the end
			if i-j > 1 {
				// compute count
				c := i - j

				// check when end of array add 1 to count to cater for non moved pointer
				if i == len(chars)-1 {
					c++
				}

				result = result + strconv.Itoa(c)
			}

			// update j to point to i
			j = i
		}
	}

	return result
}
