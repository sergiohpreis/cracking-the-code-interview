package arraystrings

import (
	"strings"
)

/*
	--------------------------------
	# Hint 106
	You do not have to - and should not - generate all
	permutations. This would be very inefficient
	--------------------------------
	# Hint 121
	What characteristics would a string that is a permutation
	of a palindrome have?
	--------------------------------
	# Hint 134
	Have you tried a hash table? You should be able to get this
	down to O(N) time
	--------------------------------
	# Hint 136
	Can you reduce the space usage by using a bit vector?
	--------------------------------
*/

/*
My implementation

Time Complexity : O(N)
Space Complexity: O(1)
*/
func PalindromePermutation(s string) bool {
	var sum int
	ocurrences := make(map[string]int)

	for _, c := range strings.ToLower(s) {
		char := string(c)

		if char != " " {
			ocurrences[char]++
		}

		if ocurrences[char] > 2 {
			ocurrences[char] = 1
		}
	}

	for _, o := range ocurrences {
		if o == 1 {
			sum++
		}
	}

	return sum <= 1
}

/*
Improved my solution using Solution #2 from the book
Here I don't need to run through the map
But still is O(N)
*/
func PalindromePermutationCountingOdds(s string) bool {
	var odd int
	ocurrences := make(map[string]int)

	for _, c := range strings.ToLower(s) {
		char := string(c)

		if char != " " {
			ocurrences[char]++

			if ocurrences[char]%2 == 1 {
				odd++
			} else {
				odd--
			}
		}
	}

	return odd <= 1
}

/*
Solution using bit vector

1. We don't need to know the counts, just if is even or odd

We can use a single int as a bit vector, left shifting 1 to
the rune
After we can toggle the bit at that value (if even will
stay at 0)

In the end, we check if only one bit is set to 1 by comparing
the int with 0

Example #1: 0001000
If we subtract 1 from this number: 00001111
00010000 &
00001111
--------
000000000

Example #2: 00000100
If we subtract 1 from this number: 00000011
00000100 &
00000011
--------
00000000

Example #3: 00101000
If we subtract 1 from this number: 00100111
00000100 &
00100111
--------
00000100

If we subtract 1 from the number and compare with the number
before subtraction, if only have 1 bit setted, will return 0

(bitVector & (bitVector - 1)) == 0
-> will return true if only 1 bit setted
*/
func createBitVector(s string) int {
	var bitVector int

	for _, c := range strings.ToLower(s) {
		bitVector = toggle(bitVector, c)
	}
	return bitVector
}

func toggle(bitVector int, index rune) int {
	if index < 0 {
		return bitVector
	}

	mask := 1 << index
	if (bitVector & mask) == 0 {
		bitVector |= mask
	} else {
		bitVector &= ^mask
	}

	return bitVector
}

func PalindromePermutationUsingBitVector(s string) bool {
	bitVector := createBitVector(s)
	return bitVector == 0 || checkExactlyOneBitSet(bitVector)
}

func checkExactlyOneBitSet(bitVector int) bool {
	return (bitVector & (bitVector - 1)) == 0
}
