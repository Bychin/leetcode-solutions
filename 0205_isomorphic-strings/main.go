package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	forwardMap := make(map[byte]byte, len(s))
	backwardMap := make(map[byte]byte, len(s))

	for i := 0; i < len(s); i++ {
		char1, ok1 := forwardMap[s[i]]
		char2, ok2 := backwardMap[t[i]]
		if !ok1 {
			if ok2 {
				return false
			}

			forwardMap[s[i]] = t[i]
			backwardMap[t[i]] = s[i]
			continue
		}

		if char1 != t[i] || char2 != s[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
}
