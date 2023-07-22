package main

import (
	"fmt"
)

func equalCounts(s, p map[byte]int) bool {
	if len(s) != len(p) {
		return false
	}

	for k, v := range p {
		if s[k] != v {
			return false
		}
	}

	return true
}

func findAnagrams(s string, p string) (result []int) {
	if len(s) < len(p) {
		return
	}

	pCount := make(map[byte]int, 26)
	windowCount := make(map[byte]int, 26)

	for i := 0; i < len(p); i++ {
		pCount[p[i]] = pCount[p[i]] + 1
		windowCount[s[i]] = windowCount[s[i]] + 1
	}

	left := 0
	right := len(p) - 1

	for right < len(s) {
		if left != 0 {
			// delete previous left symbol
			val := windowCount[s[left-1]]
			windowCount[s[left-1]] = val - 1
			if val-1 == 0 {
				delete(windowCount, s[left-1])
			}

			// add new right symbol
			val = windowCount[s[right]]
			windowCount[s[right]] = val + 1
		}

		// compare current window counter & p counter
		if equalCounts(windowCount, pCount) {
			result = append(result, left)
		}

		left++
		right++
	}

	return
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
	fmt.Println(findAnagrams("aa", "bb"))
}
