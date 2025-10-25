package main

import (
	"fmt"
)

func uniqueOccurrences(arr []int) bool {
	counts := map[int]int{}
	for _, n := range arr {
		counts[n]++
	}

	uniqueCounts := map[int]struct{}{}
	for _, count := range counts {
		if _, ok := uniqueCounts[count]; ok {
			return false
		}
		uniqueCounts[count] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))                 // true
	fmt.Println(uniqueOccurrences([]int{1, 2}))                             // false
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0})) // true
}
