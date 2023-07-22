package main

import "fmt"

func findMin(nums []int) int {
	min := 0
	max := len(nums) - 1

	for min <= max {
		pivot := (max + min) / 2

		if pivot > 0 {
			if nums[pivot-1] >= nums[pivot] {
				return nums[pivot]
			}
		} else {
			if nums[len(nums)-1] >= nums[pivot] {
				return nums[pivot]
			}
		}

		if pivot < len(nums)-1 {
			if nums[pivot] >= nums[pivot+1] {
				return nums[pivot+1]
			}
		} else {
			if nums[pivot] >= nums[0] {
				return nums[0]
			}
		}

		if nums[pivot] > nums[len(nums)-1] {
			min = pivot
			continue
		}
		if nums[pivot] < nums[0] {
			max = pivot
			continue
		}

		return nums[0]
	}

	return -1
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{1, 2, 3, 4, 5}))
	fmt.Println(findMin([]int{10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
