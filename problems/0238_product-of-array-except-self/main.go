package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = 1
			continue
		}
		result[i] = result[i-1] * nums[i-1]
	}

	tmp := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			continue
		}
		tmp *= nums[i+1]
		result[i] *= tmp
	}

	return result
}

var testCases = [][]int{
	{1, 2, 3, 4},      // [24,12,8,6]
	{-1, 1, 0, -3, 3}, // [0,0,9,0,0]
}

func main() {
	for _, tc := range testCases {
		fmt.Println(productExceptSelf(tc))
	}
}
