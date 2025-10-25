package main

import "fmt"

func minOperations(nums1 []int, nums2 []int) int64 {
	operations := 1 // for append

	x := nums2[len(nums2)-1]
	minDiffForX := 10 << 20
	lucky := false // x is between nums1[i] and nums2[i] for some i

	for i := 0; i < len(nums1); i++ {
		if nums1[i] <= x && x <= nums2[i] {
			lucky = true
			break
		}
		if nums1[i] >= x && x >= nums2[i] {
			lucky = true
			break
		}
	}

	for i := 0; i < len(nums1); i++ {
		if nums1[i] < nums2[i] {
			operations += nums2[i] - nums1[i]
		} else {
			operations += nums1[i] - nums2[i]
		}
		if lucky {
			continue
		}
		a, b := nums1[i], nums2[i]
		if nums1[i] > nums2[i] {
			b, a = a, b
		}
		if x < a {
			minDiffForX = min(minDiffForX, a-x)
		}
		if b < x {
			minDiffForX = min(minDiffForX, x-b)
		}
	}

	if !lucky {
		operations += minDiffForX
	}
	return int64(operations)
}

func main() {
	fmt.Println(minOperations([]int{2, 8}, []int{1, 7, 3}))           // 4
	fmt.Println(minOperations([]int{1, 3, 6}, []int{2, 4, 5, 3}))     // 4
	fmt.Println(minOperations([]int{2}, []int{3, 4}))                 // 3
	fmt.Println(minOperations([]int{1, 2}, []int{1, 3, 4}))           // 3
	fmt.Println(minOperations([]int{2}, []int{3, 5}))                 // 4
	fmt.Println(minOperations([]int{753, 357}, []int{271, 520, 413})) // 646
}
