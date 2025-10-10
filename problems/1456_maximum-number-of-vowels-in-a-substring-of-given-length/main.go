package main

import (
	"fmt"
)

func maxVowels(s string, k int) int {
	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	i, j := 0, 0
	maxCount := 0
	currCount := 0

	for j < len(s) {
		if j-i+1 <= k {
			if _, ok := vowels[s[j]]; ok {
				currCount++
			}
			j++
		} else {
			maxCount = max(maxCount, currCount)
			if _, ok := vowels[s[i]]; ok {
				currCount--
			}
			i++
		}
	}

	return max(maxCount, currCount)
}

func main() {
	fmt.Println(maxVowels("abciiidef", 3)) // 3
	fmt.Println(maxVowels("aeiou", 2))     // 2
	fmt.Println(maxVowels("leetcode", 3))  // 2
}
