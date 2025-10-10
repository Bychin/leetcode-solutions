package main

import (
	"fmt"
)

// prefix sum

type NumArray struct {
	prefixSum map[int]int
}

func Constructor(nums []int) NumArray {
	prefixSum := make(map[int]int, len(nums))
	cumSum := 0
	for i := 0; i < len(nums); i++ {
		cumSum += nums[i]
		prefixSum[i] = cumSum
	}
	return NumArray{prefixSum: prefixSum}
}

func (n *NumArray) SumRange(left int, right int) int {
	return n.prefixSum[right] - n.prefixSum[left-1]
}

func main() {
	n := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(n.SumRange(0, 2)) // 1
	fmt.Println(n.SumRange(2, 5)) // -1
	fmt.Println(n.SumRange(0, 5)) // -3
	fmt.Println(n.SumRange(1, 5)) // -1
	fmt.Println(n.SumRange(4, 5)) // 1
}
