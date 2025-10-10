package main

import (
	"fmt"
)

func sumOfUnique(nums []int) int {
	count := make(map[int]int, len(nums))
	sum := 0

	for _, n := range nums {
		if c, ok := count[n]; !ok {
			sum += n
			count[n] = 1
			continue
		} else {
			if c == 1 {
				sum -= n
			}
			count[n] = c + 1
		}
	}

	return sum
}

func main() {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 2}))    // 4
	fmt.Println(sumOfUnique([]int{1, 1, 1, 1, 1})) // 0
	fmt.Println(sumOfUnique([]int{1, 2, 3, 4, 5})) // 15
}
