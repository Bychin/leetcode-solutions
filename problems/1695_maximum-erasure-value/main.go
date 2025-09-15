package main

import (
	"fmt"
)

func maximumUniqueSubarray(nums []int) int {
	set := map[int]struct{}{}
	currScore, maxScore := 0, 0

	i, j := 0, 0
	for j < len(nums) {
		if _, ok := set[nums[j]]; !ok {
			set[nums[j]] = struct{}{}
			currScore += nums[j]
			j++
			continue
		}
		// nums[j] was already found, move left side of 'window'
		if currScore > maxScore {
			maxScore = currScore
		}
		for ; i < j; i++ {
			currScore -= nums[i]
			delete(set, nums[i])
			if nums[i] == nums[j] {
				i++
				break
			}
		}
	}

	if currScore > maxScore {
		return currScore
	}
	return maxScore
}

func main() {
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}) == 17)
	fmt.Println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5}) == 8)
	fmt.Println(maximumUniqueSubarray([]int{1, 2, 3, 4, 5, 1, 2, 3, 4}) == 15)
}
