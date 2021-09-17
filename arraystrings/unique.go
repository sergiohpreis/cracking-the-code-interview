package arraystrings

/*
	--------------------------------
	# Hint 44
	Try a Hash Table
	--------------------------------
	# Hint 117
	Could a bit vector be useful?
	R: We can reduce our space usage by a factor of eight by using a bit vector. We will assume,
	in the below code that the string only uses the lowercase
	letters a through z. This will allowus to use just a
	single int
	--------------------------------
	# Hint 132
	Can you solve it in O(N log N) time? What might
	a solution like that look like?
	R: If we cannot use additional data structures, we could:

	1. Compare every char of the string, to every other char of the string,
	using a nested loop
	Time Complexity: O(N*N)

	2. Sort the string - O(N log N)
	and than compare every char with the neighboor char  - O(N)
	Time Complexity: O(N log N)
	Obs: Sort Algorithms take extra space
	--------------------------------
*/

/*
	My Solution

	Time Complexity  : O(N)
	Space Complexity : O(1)
*/
func IsUnique(s string) bool {
	occurrences := make(map[string]int)

	for _, char := range s {

		if occurrences[string(char)] > 0 {
			return false
		} else {
			occurrences[string(char)] += 1
		}
	}

	return true
}

/*
	Solution from the book

	Time Complexity  : O(N)
	Space Complexity : O(1)
	Where N is the length of the string

	The time complexity can be expressed as O(1) since the for loop
	will never iterate through more than 128 characters (ASCII)

	If characters set is not fixed, can be expressed as:
	Time Complexity  : O(min(C,N)) or 0(C)
	Space Complexity : O(C)
	Where C is the size of the character set
*/
func IsUniqueUsingBoolArray(s string) bool {
	if len(s) > 128 {
		return false
	}

	var char_set [128]bool

	for _, char := range s {
		if char_set[char] {
			return false
		}
		char_set[char] = true
	}

	return true
}

/*
Solution from the book, using a bit vector

1. Reduce the space usage by a factor of 8
2. Assuming that strings only uses lowercase
letters a through z (single int)
*/
func IsUniqueUsingBitVector(s string) bool {
	checker := 0
	for _, char := range s {
		val := char - rune('a')
		if (checker & (1 << val)) > 0 {
			return false
		}
		checker |= (1 << val)
	}

	return true
}
