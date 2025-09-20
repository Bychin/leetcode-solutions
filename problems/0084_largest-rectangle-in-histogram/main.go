package main

import (
	"fmt"
)

type pair struct {
	index  int
	height int
}

func largestRectangleArea(heights []int) int {
	stack := []pair{}
	maxArea := 0

	for i, h := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1].height > h {
			index := stack[len(stack)-1].index
			height := stack[len(stack)-1].height
			stack = stack[:len(stack)-1]

			area := (i - index) * height
			maxArea = max(maxArea, area)

			start = index
		}
		stack = append(stack, pair{
			index:  start,
			height: h,
		})
	}

	i := len(heights)
	for len(stack) > 0 {
		index := stack[len(stack)-1].index
		height := stack[len(stack)-1].height
		stack = stack[:len(stack)-1]

		area := (i - index) * height
		maxArea = max(maxArea, area)
	}

	return maxArea
}

func main() {
	fmt.Println(largestRectangleArea([]int{5, 3, 4, 5, 3, 5}) == 18)
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}) == 10)
	fmt.Println(largestRectangleArea([]int{2, 4}) == 4)
	fmt.Println(largestRectangleArea([]int{0}) == 0)
}
