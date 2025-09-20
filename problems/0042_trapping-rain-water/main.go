package main

import (
	"fmt"
)

// two pointers

func trap(height []int) (result int) {
	i, j := 0, len(height)-1
	currHeight := 0
	for i < j {
		if height[i] < height[j] {
			if height[i] <= currHeight {
				result += currHeight - height[i]
			} else {
				currHeight = min(height[i], height[j])
			}
			i++
			continue
		} else {
			if height[j] <= currHeight {
				result += currHeight - height[j]
			} else {
				currHeight = min(height[i], height[j])
			}
			j--
			continue
		}
	}
	return result
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   // 9
	fmt.Println(trap([]int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1}))       // 9
}
