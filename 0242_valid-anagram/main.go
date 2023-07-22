package main

import (
	"fmt"
)

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
