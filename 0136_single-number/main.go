package main

import (
	"fmt"
)

func singleNumber(nums []int) (result int) {
	for _, n := range nums {
		result ^= n
	}

	return result
}

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{1}))
}
