package main

import "fmt"

func findMin(nums []int) int {
	min := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	return min
}

func minMoves(nums []int) (moves int) {
	min := findMin(nums)

	for _, n := range nums {
		moves += (n - min)
	}

	return
}

func main() {
	fmt.Println(minMoves([]int{1, 2, 3}))
	fmt.Println(minMoves([]int{1, 1, 1}))
	fmt.Println(minMoves([]int{10, 1, 5}))
	fmt.Println(minMoves([]int{1, 1, 1000000000}))
	fmt.Println(minMoves([]int{5, 6, 8, 8, 5}))
	fmt.Println(minMoves([]int{5, 5, 6, 8, 8}))
}
