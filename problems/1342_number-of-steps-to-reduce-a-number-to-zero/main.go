package main

import (
	"fmt"
)

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}

	bitsLen := 0
	totalLen := 0

	for num != 0 {
		if num&1 == 1 {
			bitsLen++
		}
		totalLen++
		num >>= 1
	}

	return totalLen - 1 + bitsLen
}

func main() {
	fmt.Println(numberOfSteps(14))
}
