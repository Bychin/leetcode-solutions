package main

import (
	"fmt"
)

func numIdenticalPairs(nums []int) int {
	counts := map[int]int{}
	result := 0

	for _, n := range nums {
		result += counts[n]
		counts[n]++
	}

	return result
}

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})) // 4
	fmt.Println(numIdenticalPairs([]int{1, 1, 1, 1}))       // 6
	fmt.Println(numIdenticalPairs([]int{1, 2, 3}))          // 0
}
