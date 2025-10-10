package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[slow] == 0 && nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
		if nums[slow] != 0 {
			slow++
		}
	}
}

func main() {
	fmt.Println(func() []int { a := []int{0, 1, 0, 3, 12}; moveZeroes(a); return a }()) // [1,3,12,0,0]
	fmt.Println(func() []int { a := []int{1, 2, 0, 3, 12}; moveZeroes(a); return a }()) // [1,2,3,12,0]
	fmt.Println(func() []int { a := []int{0}; moveZeroes(a); return a }())              // [0]
}
