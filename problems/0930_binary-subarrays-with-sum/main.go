package main

import "fmt"

// PrefixSum: O(n) time and O(n) space
func numSubarraysUsingPrefixSum(nums []int, goal int) (result int) {
	sum := 0
	freqs := map[int]int{}

	for _, n := range nums {
		sum += n

		if sum == goal {
			result++
		}

		diff := sum - goal
		if f, ok := freqs[diff]; ok {
			result += f
		}
		freqs[sum]++
	}

	return result
}

// Sliding Window: O(n) time and O(1) space
// Because we have zeroes in nums, using SW we can calculate
// the amount of subarrs with sum <= goal and not ==.
// Thus, the result will be (sum <= goal) - (sum <= goal-1).
func numSubarraysWithSum(nums []int, goal int) (result int) {
	return helperSumAtMost(nums, goal) - helperSumAtMost(nums, goal-1)
}

func helperSumAtMost(nums []int, goal int) (result int) {
	sum := 0
	for i, j := 0, 0; j < len(nums); j++ {
		sum += nums[j]
		for sum > goal && i <= j {
			sum -= nums[i]
			i++
		}
		result += j - i + 1
	}
	return result
}

func main() {
	fmt.Println(numSubarraysUsingPrefixSum([]int{1, 0, 1, 0, 1}, 2)) // 4
	fmt.Println(numSubarraysUsingPrefixSum([]int{0, 0, 0, 0, 0}, 0)) // 15

	fmt.Println(numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2)) // 4
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0)) // 15
}
