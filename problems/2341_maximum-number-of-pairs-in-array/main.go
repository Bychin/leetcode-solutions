package main

import (
	"fmt"
)

func numberOfPairs(nums []int) []int {
	freq := make(map[int]int, len(nums))
	for _, num := range nums {
		freq[num] += 1
	}

	result := make([]int, 2)

	for _, v := range freq {
		result[0] += v / 2
		result[1] += v % 2
	}

	return result
}

func main() {
	fmt.Println(numberOfPairs([]int{1, 3, 2, 1, 3, 2, 2}))
}
