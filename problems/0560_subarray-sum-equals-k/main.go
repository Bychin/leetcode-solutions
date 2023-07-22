package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) (result int) {
	prefixSumCount := map[int]int{0: 1}

	cumSum := 0
	for _, n := range nums {
		cumSum += n

		result += prefixSumCount[cumSum-k]
		prefixSumCount[cumSum] += 1
	}

	return
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum([]int{1, -1, 1}, 1))
}
