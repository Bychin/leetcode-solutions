package main

import (
	"fmt"
)

func minArrivalsToDiscard(arrivals []int, w int, m int) int {
	result := 0

	counts := make(map[int]int)
	for i := 0; i < w; i++ {
		item := arrivals[i]
		if counts[item] >= m {
			arrivals[i] = -1
			result++
		} else {
			counts[item]++
		}
	}

	leftBoundary := 0
	for i := w; i < len(arrivals); i++ {
		if arrivals[leftBoundary] != -1 {
			outItem := arrivals[leftBoundary]
			counts[outItem]--
		}
		leftBoundary++

		inItem := arrivals[i]
		if counts[inItem] >= m {
			arrivals[i] = -1
			result++
		} else {
			counts[inItem]++
		}
	}

	return result
}

func main() {
	fmt.Println(minArrivalsToDiscard([]int{1, 2, 1, 3, 1}, 4, 2))                                                                              // Output: 0
	fmt.Println(minArrivalsToDiscard([]int{1, 2, 3, 3, 3, 4}, 3, 2))                                                                           // Output: 1
	fmt.Println(minArrivalsToDiscard([]int{3, 3, 3, 3, 3, 4}, 3, 2))                                                                           // Output: 1
	fmt.Println(minArrivalsToDiscard([]int{10, 4, 3, 6, 4, 5, 6, 1, 4}, 7, 1))                                                                 // Output: 2
	fmt.Println(minArrivalsToDiscard([]int{7, 3, 9, 9, 7, 3, 5, 9, 7, 2, 6, 10, 9, 7, 9, 1, 3, 6, 2, 4, 6, 2, 6, 8, 4, 8, 2, 7, 5, 6}, 10, 1)) // Output: 13
}
