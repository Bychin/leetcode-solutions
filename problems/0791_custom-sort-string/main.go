package main

import (
	"fmt"
)

func customSortString(order string, s string) string {
	counts := map[byte]int{}
	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}

	result := []byte{}
	for i := 0; i < len(order); i++ {
		char := order[i]
		count, ok := counts[char]
		if !ok {
			continue
		}
		for j := 0; j < count; j++ {
			result = append(result, char)
		}
		delete(counts, char)
	}
	for char, count := range counts {
		for j := 0; j < count; j++ {
			result = append(result, char)
		}
	}
	return string(result)
}

func main() {
	fmt.Println(customSortString("cba", "abcd"))   // "cbad"
	fmt.Println(customSortString("bcafg", "abcd")) // "bcad"
}
