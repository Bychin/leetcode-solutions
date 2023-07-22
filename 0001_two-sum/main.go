package main

import (
	"fmt"
)

// speed: O(n^2), memory: O(1)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1} // unreachable code
}

// speed: O(n), memory: O(n)
func twoSum2(nums []int, target int) []int {
	lookupTable := map[int]int{}

	for i := len(nums) - 1; i >= 0; i-- {
		diff := target - nums[i]
		j, ok := lookupTable[diff]
		if !ok {
			lookupTable[nums[i]] = i
			continue
		}

		return []int{i, j}
	}

	return []int{-1, -1} // unreachable code
}

func main() {
	funcToCall := twoSum2

	fmt.Println(funcToCall([]int{-1, -2, -3, -4, -5}, -8))
	fmt.Println(funcToCall([]int{2, 7, 11, 15}, 9))
	fmt.Println(funcToCall([]int{2, 5, 2, 5}, 7))
}
