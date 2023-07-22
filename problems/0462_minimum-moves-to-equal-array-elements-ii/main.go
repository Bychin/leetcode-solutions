package main

import (
	"fmt"
	"sort"
)

func findMedian(sortedNums []int) int {
	if len(sortedNums)%2 == 0 {
		return (sortedNums[len(sortedNums)/2] + sortedNums[(len(sortedNums)/2)-1]) / 2
	}

	return sortedNums[len(sortedNums)/2]
}

func minMoves2(nums []int) (result int) {
	sort.Ints(nums)
	median := findMedian(nums)

	for _, n := range nums {
		if median > n {
			result += median - n
		} else {
			result += n - median
		}
	}

	return
}

func main() {
	fmt.Println(minMoves2([]int{1, 2, 3}))
	fmt.Println(minMoves2([]int{1, 1, 1}))
	fmt.Println(minMoves2([]int{1, 10, 2, 9}))
	fmt.Println(minMoves2([]int{1, 1, 1000000000}))
	fmt.Println(minMoves2([]int{5, 6, 8, 8, 5}))
	fmt.Println(minMoves2([]int{5, 5, 6, 8, 8}))
}
