package main

import (
	"fmt"
	"strings"
)

func frequencySort(s string) string {
	counts := map[rune]int{}
	maxCount := 0 // it is always <= len(s)

	for _, r := range s {
		counts[r]++
		maxCount = max(maxCount, counts[r])
	}

	buckets := make([][]rune, maxCount+1)
	for r, count := range counts {
		buckets[count] = append(buckets[count], r)
	}

	result := ""
	for i := maxCount; i > 0; i-- {
		for _, r := range buckets[i] {
			result += strings.Repeat(string(r), i)
		}
	}

	return result
}

func main() {
	fmt.Println(frequencySort("tree"))   // "eert"
	fmt.Println(frequencySort("cccaaa")) // "aaaccc"
	fmt.Println(frequencySort("Aabb"))   // "bbAa"
}
