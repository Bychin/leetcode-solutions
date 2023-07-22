package main

import (
	"fmt"
)

func pivotIndex(nums []int) int {
	totalSum := 0
	for _, n := range nums {
		totalSum += n
	}

	leftSum := 0
	rightSum := totalSum

	if leftSum == rightSum-nums[0] {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		leftSum += nums[i-1]
		rightSum -= nums[i-1]
		if leftSum == rightSum-nums[i] {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
	fmt.Println(pivotIndex([]int{1, -1, 2}))
}
