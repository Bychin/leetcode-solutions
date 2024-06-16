package main

import (
	"fmt"
)

// 2 solutions:
// 1. sort + check neighbors	time: O(nlogn),	memory: O(1)
// 2. set + check on insert		time: O(n),		memory: O(n)

func containsDuplicate(nums []int) bool {
	duplicates := make(map[int]struct{}, len(nums))

	for _, n := range nums {
		if _, ok := duplicates[n]; ok {
			return true
		}

		duplicates[n] = struct{}{}
	}

	return false
}

var testCases = [][]int{
	{1, 7, 3, 6, 5, 6},
	{1, 2, 3},
	{2, 1, -1},
	{1, -1, 2},
}

func main() {
	for _, tc := range testCases {
		fmt.Println(containsDuplicate(tc))
	}
}
