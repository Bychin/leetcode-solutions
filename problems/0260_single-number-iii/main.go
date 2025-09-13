package main

import "fmt"

// 1. use XOR to find the result of XOR of two numbers that we are looking for
// 2. find bit with value '1' in the result - group all numbers with bit '0' in one group and with bit '1' in another
// 3. use XOR on both groups to find two numbers that we are looking for

func singleNumber(nums []int) []int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}

	mask := 0 // mask with one bit set on position where two unique numbers differs
	for bitNum := 0; ; bitNum++ {
		mask = 1 << bitNum
		if xor&mask != 0 {
			break // found
		}
	}

	xor0, xor1 := 0, 0
	for _, n := range nums {
		if n&mask == 0 {
			xor0 ^= n
		} else {
			xor1 ^= n
		}
	}

	return []int{xor0, xor1}
}

var testCases = [][]int{
	{1, 2, 1, 3, 2, 5}, // [5, 3]
	{-1, 0},            // [0, -1]
	{0, 1},             // [0, 1]
}

func main() {
	for _, tc := range testCases {
		fmt.Println(singleNumber(tc))
	}
}
