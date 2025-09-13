package main

import (
	"fmt"
	"slices"
)

func plusOne(digits []int) []int {
	carry := false

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 <= 9 {
			digits[i] += 1
			return digits
		}

		digits[i] = 0
		carry = true
		continue
	}

	if carry {
		digits = slices.Insert(digits, 0, 1)
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3})) // [1,2,4]
	fmt.Println(plusOne([]int{9}))       // [1,0]
}
