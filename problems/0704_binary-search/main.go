package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	min := 0
	max := len(nums) - 1

	for min <= max {
		pivot := (max + min) / 2

		if nums[pivot] == target {
			return pivot
		}
		if target < nums[pivot] {
			max = pivot - 1
		}
		if target > nums[pivot] {
			min = pivot + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println(search([]int{-1}, -1))
	fmt.Println(search([]int{-1}, 2))
	fmt.Println(search([]int{5}, 5))
}
