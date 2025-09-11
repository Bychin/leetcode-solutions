package main

import (
	"fmt"
)

func checkUpwards(lookup map[int]struct{}, value, counter int) int {
	if _, ok := lookup[value+1]; !ok {
		return counter
	}
	return checkUpwards(lookup, value+1, counter+1)
}

func longestConsecutive(nums []int) int {
	lookup := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		lookup[n] = struct{}{}
	}

	maxCounter := 0
	for n := range lookup {
		if _, ok := lookup[n-1]; ok {
			// not the start of a sequence
			continue
		}
		currentCounter := checkUpwards(lookup, n-1, 0)
		if currentCounter > maxCounter {
			maxCounter = currentCounter
		}
	}

	return maxCounter
}

var testCases = [][]int{
	{100, 4, 200, 1, 3, 2},         // 4
	{100, 4, 101, 1, 3, 2},         // 4
	{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, // 9
}

func main() {
	for _, tc := range testCases {
		fmt.Println(longestConsecutive(tc))
	}
}
