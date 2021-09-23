package arraystrings

import (
	"fmt"
	"strings"
)

/*
	--------------------------------
	Hint #92
	Do the easy thing first. Compress the string, then compare
	the lengths.
	--------------------------------
	Hint #110
	Be careful that you aren't repeatedly concatenating strings
	together. This can be very inefficient
	--------------------------------
*/

/*
My solution
Time  Complexity: O(N)
Space Complexity: O(N)
*/

// TODO: Search if regex is better than loop (time complexity)
func StringCompression(s string) string {
	var counter int
	var compressed []string
	var currentChar string

	for _, c := range s {
		if currentChar == "" {
			currentChar = string(c)
			counter = 1
		} else if currentChar != string(c) {
			compressed = append(compressed, currentChar, fmt.Sprint(counter))
			currentChar = string(c)
			counter = 1
		} else {
			counter++
		}
	}

	if currentChar != compressed[len(compressed)-1] {
		compressed = append(compressed, currentChar, fmt.Sprint(counter))
	}

	if len(compressed) > len(s) {
		return s
	}

	return strings.Join(compressed, "")
}

/*
Book bad implementation
Time Complexity: O (p + k*k)
Where P is the size of original string
Where K is the number of character sequences
*/
func CompressBad(str string) string {
	var compressedString string
	var countConsecutive int

	for i := 0; i < len(str); i++ {
		countConsecutive++

		if i+1 >= len(str) || str[i] != str[i+1] {
			compressedString += string(str[i]) + fmt.Sprint(countConsecutive)
			countConsecutive = 0
		}
	}

	if len(compressedString) > len(str) {
		return str
	}

	return compressedString
}

// Book using string builder
func CompressStringBuilder(str string) string {
	var compressed []string
	var countConsecutive int

	for i := 0; i < len(str); i++ {
		countConsecutive++

		if i+1 >= len(str) || str[i] != str[i+1] {
			compressed = append(compressed, string(str[i]), fmt.Sprint(countConsecutive))
			countConsecutive = 0
		}
	}

	if len(compressed) > len(str) {
		return str
	}

	return strings.Join(compressed, "")
}

func countCompression(str string) int {
	var compressedLength, countConsecutive int

	for i := 0; i < len(str); i++ {
		countConsecutive++

		if i+1 >= len(str) || str[i] != str[i+1] {
			compressedLength += 1 + len(fmt.Sprint(countConsecutive))
			countConsecutive = 0
		}
	}

	return compressedLength
}

/*
Book solution counting compressed size

Tradeoffs:
1. One more loop needed, but we do not create a string that we
will not use
2. []string created with the necessarie capacity avoiding
to double the size
*/
func CompressCounting(str string) string {
	finalLength := countCompression(str)

	if finalLength > len(str) {
		return str
	}

	compressed := make([]string, finalLength)
	var countConsecutive int

	for i := 0; i < len(str); i++ {
		countConsecutive++

		if i+1 >= len(str) || str[i] != str[i+1] {
			compressed = append(compressed, string(str[i]), fmt.Sprint(countConsecutive))
			countConsecutive = 0
		}
	}

	if len(compressed) > len(str) {
		return str
	}

	return strings.Join(compressed, "")
}
