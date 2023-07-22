package main

import (
	"fmt"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	topsByWord := make(map[string]int, len(words))
	uniqueWords := make([]string, len(words))

	for _, word := range words {
		if _, ok := topsByWord[word]; !ok {
			uniqueWords = append(uniqueWords, word)
		}

		topsByWord[word]++
	}

	sort.Slice(uniqueWords, func(i, j int) bool {
		if topsByWord[uniqueWords[i]] == topsByWord[uniqueWords[j]] {
			return uniqueWords[i] < uniqueWords[j]
		}

		return topsByWord[uniqueWords[i]] > topsByWord[uniqueWords[j]]
	})

	return uniqueWords[:k]
}

func main() {
	fmt.Println(
		topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2),
	)
}
