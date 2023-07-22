package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	result := make([]int, 0, len(nums))

	fmt.Println(nums)

	// prefix
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result = append(result, nums[i])
			continue
		}

		val := result[len(result)-1] * nums[i]
		result = append(result, val)
	}

	fmt.Println(result)

	// postfix
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			result[i] = result[i-1]
			continue
		}

		result[i] *= nums[i+1]
	}

	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))      // [24,12,8,6]
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3})) // [0,0,9,0,0]
}
