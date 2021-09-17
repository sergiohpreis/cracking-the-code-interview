package main

import (
	"fmt"
	"strings"

	"github.com/sergiohpreis/crackingcodeinterview/arraystrings"
)

func isUnique(as []string) {
	for _, s := range as {
		fmt.Printf("\"%v\" is unique?\n", s)
		fmt.Println("My implementation ->", arraystrings.IsUnique(s))
		fmt.Println("Using Bool Array  ->", arraystrings.IsUniqueUsingBoolArray(s))
		fmt.Println("Using Bit Vector  ->", arraystrings.IsUniqueUsingBitVector(s))
		fmt.Println()
	}
	fmt.Println()
}

func permutation(aas [][]string) {
	for _, as := range aas {
		fmt.Printf("\"%v\" is a permutation of \"%v\"?\n", as[0], as[1])
		fmt.Println("My implementation using sort    ->", arraystrings.IsPermutationUsingSort(as[0], as[1]))
		fmt.Println("My implementation summing bytes ->", arraystrings.IsPermutationSummingBytes(as[0], as[1]))
		fmt.Println("My implementation using hash    ->", arraystrings.IsPermutationUsingHashTable(as[0], as[1]))
		fmt.Println("Book implementation using array ->", arraystrings.IsPermutationUsingArray(as[0], as[1]))
		fmt.Println()
	}
	fmt.Println()
}

func URLify(m map[string]int) {
	for s, l := range m {
		fmt.Printf("URLify the string: \"%v\" (length: %v)\n", s, l)
		fmt.Println("My implementation using loop and string slice ->", arraystrings.URLifyUsingSplit(s, l))
		fmt.Println("My implementation using regex                 ->", arraystrings.URLifyUsingRegex(s, l))
		ar := strings.Split(s, "")
		fmt.Println("Book implementation counting spaces           ->", arraystrings.URLifyCountingSpaces(ar, l))
		fmt.Println()
	}
	fmt.Println()
}

func PalindromePermutation(as []string) {
	for _, s := range as {
		fmt.Printf("The string \"%v\" is a palindrome permutation?\n", s)
		fmt.Println("My implementation using a map         ->", arraystrings.PalindromePermutation(s))
		fmt.Println("My implementation counting odds       ->", arraystrings.PalindromePermutationCountingOdds(s))
		fmt.Println("Book implementation using bit vector  ->", arraystrings.PalindromePermutationUsingBitVector(s))
		fmt.Println()
	}
}

func arrayStrings() {
	isUnique([]string{
		"abcd",
		"aabc",
		"defg",
	})
	permutation([][]string{
		{"god     ", "dog"},
		{"god     ", "dog     "},
		{"Dog", "dog"},
		{"abbc", "bbad"},
		{"abc", "bca"},
		{"abc", "cbb"},
		{"abc", "abcd"},
	})
	URLify(map[string]int{
		"Mr John Smith     ": 13,
	})
	PalindromePermutation([]string{
		"Tact Coa",      // taco cat, // atc o cta -> true
		"dad",           // dad                    -> true
		"tta e t t",     //                        -> false
		"tt a t a tt t", // ttt aa ttt, tatt  ttat -> true
	})
}

func main() {
	arrayStrings()
}
