package main

import (
	"fmt"
)

func countOdds(low int, high int) (result int) {
	result = (high - low) / 2
	if high%2 != 0 || low%2 != 0 {
		result++
	}

	return
}

func main() {
	fmt.Println(countOdds(3, 7))
}
