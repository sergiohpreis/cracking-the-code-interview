package arraystrings

import (
	"regexp"
	"strings"
)

/*
	--------------------------------
	# Hint 53: It's often easiest to modify strings by going
		from the end ofthe string to the beginning
	--------------------------------
	# Hint 118: You might find you need to know the number of spaces. Can you just count them?
	--------------------------------
*/

/*
My solution looping through the string length
Time Complexity : O(N)
Space Complexity: O(N)
Where N is the string length
*/
func URLifyUsingSplit(s string, l int) string {
	var a []string

	for i := 0; i < l; i++ {
		if char := string(s[i]); char == " " {
			a = append(a, "%20")
		} else {
			a = append(a, char)
		}
	}

	return strings.Join(a, "")
}

/*
My solution using regex
Time Complexity: O(1)
Space Complexity: O(1)
*/
func URLifyUsingRegex(s string, l int) string {
	s = regexp.MustCompile(`\s{2,}`).ReplaceAllString(s, "")
	return regexp.MustCompile(`\s{1}`).ReplaceAllString(s, "%20")
}

// Implementation from the book
func URLifyCountingSpaces(str []string, trueLength int) []string {
	var spaceCount int

	for i := 0; i < trueLength; i++ {
		if char := string(str[i]); char == " " {
			spaceCount++
		}
	}
	index := trueLength + spaceCount*2

	if trueLength < len(str) {
		str[trueLength] = `\0`
	}

	for i := trueLength - 1; i >= 0; i-- {
		if str[i] == " " {
			str[index-1] = "0"
			str[index-2] = "2"
			str[index-3] = "%"
			index -= 3
		} else {
			str[index-1] = str[i]
			index--
		}
	}

	return str
}
