package main

import (
	"fmt"
)

func minimumAverageDifference(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	sumForward := 0
	sumBackward := sum

	minAvDiff := 0
	index := 0

	for i := 0; i < len(nums); i++ {
		j := len(nums) - i

		sumForward += nums[i]
		sumBackward -= nums[i]

		avDiffLeft := sumForward / (i + 1)
		avDiffRight := 0
		if j > 1 {
			avDiffRight = sumBackward / (j - 1)
		}

		avDiff := avDiffLeft - avDiffRight
		if avDiff < 0 {
			avDiff = -avDiff
		}

		if minAvDiff > avDiff || i == 0 {
			minAvDiff = avDiff
			index = i
		}
	}

	return index
}

func main() {
	fmt.Println(minimumAverageDifference([]int{2, 5, 3, 9, 5, 3}))
	fmt.Println(minimumAverageDifference([]int{0}))
}
