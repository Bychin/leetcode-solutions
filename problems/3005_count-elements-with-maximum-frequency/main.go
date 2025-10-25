package main

import (
	"fmt"
)

func maxFrequencyElements(nums []int) (result int) {
	counts := map[int]int{}
	maxCount := 0
	for _, n := range nums {
		v := counts[n]
		counts[n] = v + 1
		maxCount = max(maxCount, v+1)
	}

	for _, count := range counts {
		if count == maxCount {
			result += count
		}
	}
	return result
}

func main() {
	fmt.Println(maxFrequencyElements([]int{1, 2, 2, 3, 1, 4})) // 4
	fmt.Println(maxFrequencyElements([]int{1, 2, 3, 4, 5}))    // 5
}
