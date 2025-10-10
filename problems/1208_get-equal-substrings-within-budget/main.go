package main

import (
	"fmt"
)

func absDiff(a, b byte) int {
	if a > b {
		return int(a - b)
	}
	return int(b - a)
}

func equalSubstring(s string, t string, maxCost int) int {
	maxLen := 0
	currCost := 0
	i, j := 0, 0
	for j < len(s) {
		diffJ := absDiff(s[j], t[j])

		for currCost+diffJ > maxCost {
			maxLen = max(maxLen, j-i)
			diffI := absDiff(s[i], t[i])
			currCost -= diffI
			i++
		}
		currCost += diffJ
		j++
	}
	return max(maxLen, j-i)
}

func main() {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))                        // 3
	fmt.Println(equalSubstring("abcd", "cdef", 3))                        // 1
	fmt.Println(equalSubstring("abcd", "acde", 0))                        // 1
	fmt.Println(equalSubstring("krrgw", "zjxss", 19))                     // 2
	fmt.Println(equalSubstring("krpgjbjjznpzdfy", "nxargkbydxmsgby", 14)) // 4
	fmt.Println(equalSubstring("pxezla", "loewbi", 25))                   // 4
}
