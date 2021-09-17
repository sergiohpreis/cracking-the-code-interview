package arraystrings

import (
	"sort"
	"strings"
)

/*
	--------------------------------
	Hint #1
	Describe what it means for two strings to be permutations
	of each other.
	Now, look at that definition you provided.
	Can you check the strings against that definition?
	--------------------------------
	Hint #84
	There's one solution that is O(N log N) time.
	Another solution uses some space but is O(N) time
	--------------------------------
	Hint #122
	Could a hash table be useful?
	--------------------------------
	Hint #131
	Two strings that are permutations should have the same
	characters, but in differente orders.
	Can you make orders the same?
	--------------------------------
*/

/*
Time  Complexity: 0 (N log N)
Space Complexity: O(N)
The N will always be the same, since strings will
always be the same size
*/
func IsPermutationUsingSort(x, y string) bool {
	if len(x) != len(y) {
		return false
	}

	sx := strings.Split(x, "")
	sy := strings.Split(y, "")

	sort.Strings(sx)
	sort.Strings(sy)

	x = strings.Join(sx, "")
	y = strings.Join(sy, "")

	return x == y
}

/*
Time  Complexity: 0(N)
Space Complexity: O(2N)
The N will always be the same, since strings will
always be the same size
*/
func IsPermutationSummingBytes(x, y string) bool {
	if len(x) != len(y) {
		return false
	}

	var sumX, sumY int

	for _, char := range x {
		sumX += int(char)
	}

	for _, char := range y {
		sumY += int(char)
	}

	return sumX == sumY
}

/*
Time  Complexity: 0(N)
Space Complexity: O(1)
The N will always be the same, since strings will
always be the same size
*/
func IsPermutationUsingHashTable(x, y string) bool {
	if len(x) != len(y) {
		return false
	}

	ocurrences := make(map[rune]int)

	for _, char := range x {
		ocurrences[char]++
	}

	for _, char := range y {
		ocurrences[char]--
		if ocurrences[char] < 0 {
			return false
		}
	}

	return true
}

/*
Time  Complexity: 0(N)
Space Complexity: O(1)
The N will always be the same, since strings will
always be the same size
*/
func IsPermutationUsingArray(x, y string) bool {
	if len(x) != len(y) {
		return false
	}

	var ocurrences [128]int // ASCII

	for _, char := range x {
		ocurrences[char]++
	}

	for _, char := range y {
		ocurrences[char]--
		if ocurrences[char] < 0 {
			return false
		}
	}

	return true
}
