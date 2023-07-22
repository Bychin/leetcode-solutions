package main

import (
	"fmt"
)

func runningSum(nums []int) []int {
	sum := 0
	for i, n := range nums {
		sum += n
		nums[i] = sum
	}

	return nums
}

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
}
