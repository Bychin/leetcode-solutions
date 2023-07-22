package main

import (
	"fmt"
	"sort"
)

// TODO better solution is the one without sorting and with custom hash function
// https://ru.wikipedia.org/wiki/FNV

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string, len(strs))
	for _, word := range strs {
		wordAsBytes := []byte(word)
		sort.Slice(wordAsBytes, func(i, j int) bool {
			return wordAsBytes[i] < wordAsBytes[j]
		})

		val, ok := anagrams[string(wordAsBytes)]
		if !ok {
			anagrams[string(wordAsBytes)] = []string{word}
			continue
		}

		anagrams[string(wordAsBytes)] = append(val, word)
	}

	result := make([][]string, 0, len(anagrams))
	for _, v := range anagrams {
		result = append(result, v)
	}

	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}
