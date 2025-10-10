package main

import (
	"fmt"
)

// sliding window

func minSubArrayLen(target int, nums []int) (result int) {
	i, j := 0, 1
	currSubArrSum := nums[0]
	for i < j {
		if currSubArrSum < target {
			if j < len(nums) {
				currSubArrSum += nums[j]
				j++
			} else {
				currSubArrSum -= nums[i]
				i++
			}
		}
		if currSubArrSum >= target {
			currSubArrLen := j - i
			if currSubArrLen < result || result == 0 {
				result = currSubArrLen
			}
			currSubArrSum -= nums[i]
			i++
		}
	}
	return result
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                 // 1
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))          // 3
}
