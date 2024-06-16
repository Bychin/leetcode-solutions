package main

import (
	"fmt"
)

// 2 solutions:
// 1. sort + compare											time: O(nlogn),		memory: O(1)
// 2. map of chars, increment for 1st string decrement for 2nd	time: O(len1+len2),	memory: O(len1+len2)
// Note that for 2nd solution memory is constant in case of using any alphabet.

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	lettersCount := make(map[rune]int, len(s))

	for _, c := range s {
		counter := lettersCount[c]
		lettersCount[c] = counter + 1
	}

	for _, c := range t {
		counter, ok := lettersCount[c]
		if !ok || counter == 0 {
			return false
		}

		lettersCount[c] = counter - 1
	}

	return true
}

func main() {
	fmt.Println(isAnagram("rac", "car"))
}
