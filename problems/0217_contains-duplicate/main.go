package main

import (
	"fmt"
)

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

func main() {
	fmt.Println(containsDuplicate([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(containsDuplicate([]int{1, 2, 3}))
	fmt.Println(containsDuplicate([]int{2, 1, -1}))
	fmt.Println(containsDuplicate([]int{1, -1, 2}))
}
