package main

import (
	"fmt"
	"maps"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	counters1 := map[byte]int{}
	counters2 := map[byte]int{}

	for i := range len(word1) {
		counters1[word1[i]]++
		counters2[word2[i]]++
	}

	// 1. check that the letters are the same
	if len(counters1) != len(counters2) {
		return false
	}
	for ch := range counters1 {
		if _, ok := counters2[ch]; !ok {
			return false
		}
	}

	// 2. check that the frequencies are the same
	// (we could sort [26]byte slices, but I don't like such approach)
	frequencies1 := map[int]int{}
	frequencies2 := map[int]int{}
	for _, count := range counters1 {
		frequencies1[count]++
	}
	for _, count := range counters2 {
		frequencies2[count]++
	}
	return maps.Equal(frequencies1, frequencies2)
}

func main() {
	fmt.Println(closeStrings("abc", "bca"))       // true
	fmt.Println(closeStrings("a", "aa"))          // false
	fmt.Println(closeStrings("cabbba", "abbccc")) // true
	fmt.Println(closeStrings("cabbba", "aabbss")) // false
	fmt.Println(closeStrings("uau", "ssx"))       // false
}
