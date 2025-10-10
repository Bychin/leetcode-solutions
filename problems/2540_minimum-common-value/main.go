package main

import (
	"fmt"
)

func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0

	for {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}
		if nums1[i] < nums2[j] {
			if i == len(nums1)-1 {
				return -1
			}
			i++
			continue
		}
		if nums1[i] > nums2[j] {
			if j == len(nums2)-1 {
				return -1
			}
			j++
			continue
		}
	}
}

func main() {
	fmt.Println(getCommon([]int{1, 2, 3}, []int{2, 4}))                   // 2
	fmt.Println(getCommon([]int{1, 2, 3, 6}, []int{2, 3, 4, 5}))          // 2
	fmt.Println(getCommon([]int{-9, -8, -6, 0, 1, 2}, []int{2, 3, 4, 5})) // 2
	fmt.Println(getCommon([]int{1, 1, 2}, []int{2, 4}))                   // 2
	fmt.Println(getCommon([]int{1, 3}, []int{2, 4}))                      // -1
}
