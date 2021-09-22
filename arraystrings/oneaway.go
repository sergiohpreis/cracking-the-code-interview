package arraystrings

import (
	"math"
)

/*
	--------------------------------
	Hint #23
	Start with the easy thing. Can you check each of
	the conditions separately?
	--------------------------------
	Hint #97
	What is the relationship between the "insert character"
	option and the "remove character" option? Do these
	need to be two separate checks?
	--------------------------------
	Hint #130
	Can you do all three checks in a single pass?
	--------------------------------
*/

/*
	My solution is working only for inserts or removals
	Time Complexity: Big O(X+Y)
*/
// TODO: Try to adjust to success on replaces
func OneAway(x, y string) bool {
	if math.Abs(float64(len(x)-len(y))) > 1 {
		return false
	}

	cx := countChars(x) // Big O(X)
	cy := countChars(y) // Big O(Y)
	diff := countAbsCharsDiff(cx, cy)

	return diff < 1
}

func countChars(s string) map[rune]int {
	count := make(map[rune]int)

	for _, c := range s {
		count[c]++
	}

	return count
}

func countAbsCharsDiff(x, y map[rune]int) float64 {
	var diff int

	for char := range y {
		diff += y[char] - x[char]
	}

	return math.Abs(float64(diff))
}

/*
Book solution
Time Complexity: O(N)
Where N is the length of the shorter string
*/
func OneEditAway(first, second string) bool {
	if len(first) == len(second) {
		return oneEditReplace(first, second)
	} else if len(first)+1 == len(second) { // shorter first
		return oneEditInsert(first, second)
	} else if len(first)-1 == len(second) { // shorter first
		return oneEditInsert(second, first)
	}

	return false
}

func oneEditReplace(s1, s2 string) bool {
	var foundDifference bool

	for i, c := range s1 {
		if string(c) != string(s2[i]) {
			if foundDifference {
				return false
			}
			foundDifference = true
		}
	}

	return true
}

// s1 always shorter than s2
func oneEditInsert(s1, s2 string) bool {
	var index1, index2 int

	for index2 < len(s2) && index1 < len(s1) {
		if string(s1[index1]) != string(s2[index2]) {
			if index1 != index2 {
				return false
			}
			index2++
		} else {
			index1++
			index2++
		}
	}

	return true
}

/*
Book Solution
Time Complexity: O(N)
Where N is the length of the shorter string
*/
func OneEditAwaySingle(first, second string) bool {
	// Length checks
	if math.Abs(float64(len(first)-len(second))) > 1 {
		return false
	}

	/*
		Get shorter and longer string
		s1: shorter, s2: longer
	*/
	s1, s2 := first, second
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	var index1, index2 int
	var foundDifference bool

	for index2 < len(s2) && index1 < len(s1) {
		char1 := string(s1[index1])
		char2 := string(s2[index2])
		if char1 != char2 {
			// Ensure that this is the first difference found
			if foundDifference {
				return false
			}
			foundDifference = true

			// On replace, move shorter pointer
			if len(s1) == len(s2) {
				index1++
			}
		} else {
			// If matching, move shorter pointer
			index1++
		}
		// Always move pointer for longer string
		index2++
	}

	return true
}
