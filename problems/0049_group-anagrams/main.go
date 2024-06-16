package main

import (
	"fmt"
	"sort"
)

// Solution 1:
// Sort each string and use it as a key on hashmap to collect all anagrams. Time: O(nlogn * m), memory: O(n*m)

// Solution 2:
// Instead of sorting use custom hash function like https://ru.wikipedia.org/wiki/FNV

// Solution 3:
// If we work with an alphabet (like English), we can use [26]int to count every letter and use it as a key. Time: O(n * m (* 26)), memory: O(n*m)

func groupAnagramsUsingSort(strs []string) [][]string {
	anagrams := make(map[string][]string, len(strs))
	for _, word := range strs {
		wordAsBytes := []byte(word)
		sort.Slice(wordAsBytes, func(i, j int) bool {
			return wordAsBytes[i] < wordAsBytes[j]
		})

		anagrams[string(wordAsBytes)] = append(anagrams[string(wordAsBytes)], word)
	}

	result := make([][]string, 0, len(anagrams))
	for _, v := range anagrams {
		result = append(result, v)
	}

	return result
}

func groupAnagramsForEnglish(strs []string) [][]string {
	anagrams := make(map[[26]int][]string, len(strs))
	for _, word := range strs {
		key := [26]int{}
		for _, b := range []byte(word) {
			key[b-'a'] += 1
		}

		anagrams[key] = append(anagrams[key], word)
	}

	result := make([][]string, 0, len(anagrams))
	for _, v := range anagrams {
		result = append(result, v)
	}

	return result
}

var testCases = [][]string{
	{"eat", "tea", "tan", "ate", "nat", "bat"},
	{""},
	{"a"},
}

func main() {
	for _, tc := range testCases {
		fmt.Println(groupAnagramsForEnglish(tc))
	}
}
