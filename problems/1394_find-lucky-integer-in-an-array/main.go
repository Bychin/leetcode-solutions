package main

import (
	"fmt"
)

func findLucky(arr []int) int {
	counts := map[int]int{}
	for _, n := range arr {
		counts[n]++
	}

	maxLucky := -1
	for n, count := range counts {
		if n == count {
			maxLucky = max(maxLucky, n)
		}
	}
	return maxLucky
}

func main() {
	fmt.Println(findLucky([]int{2, 2, 3, 4}))       // 2
	fmt.Println(findLucky([]int{1, 2, 2, 3, 3, 3})) // 3
	fmt.Println(findLucky([]int{2, 2, 2, 3, 3}))    // -1
}
