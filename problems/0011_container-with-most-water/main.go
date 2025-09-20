package main

import (
	"fmt"
)

// two pointers, move pointer with smaller height

func maxArea(height []int) (result int) {
	for i, j := 0, len(height)-1; i < j; {
		area := (j - i) * min(height[i], height[j])
		result = max(result, area)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return result
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea([]int{1, 1}))                      // 1
}
