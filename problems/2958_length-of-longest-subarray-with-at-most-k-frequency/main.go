package main

import "fmt"

func maxSubarrayLength(nums []int, k int) (maxLen int) {
	counters := map[int]int{}

	for i, j := 0, 0; j < len(nums); j++ {
		counters[nums[j]]++
		for counters[nums[j]] > k {
			counters[nums[i]]--
			i++
		}
		maxLen = max(maxLen, j-i+1)
	}
	return maxLen
}

func main() {
	fmt.Println(maxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2)) // 6
	fmt.Println(maxSubarrayLength([]int{1, 2, 1, 2, 1, 2, 1, 2}, 1)) // 2
	fmt.Println(maxSubarrayLength([]int{5, 5, 5, 5, 5, 5, 5}, 4))    // 4
	fmt.Println(maxSubarrayLength([]int{1}, 1))                      // 1
}
