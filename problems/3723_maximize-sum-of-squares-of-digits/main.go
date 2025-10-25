package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maxSumOfSquares(num int, sum int) string {
	b := strings.Builder{}
	b.Grow(num)

	currSum := 0

	for i := 0; i < num; i++ {
		for j := 9; j > 0; j-- { // maximize each digit
			if currSum+j <= sum {
				b.WriteString(strconv.Itoa(j))
				currSum += j
				break
			}
		}
		if currSum == sum {
			for k := i + 1; k < num; k++ {
				b.WriteString("0")
			}
			break
		}
	}

	if currSum != sum {
		return ""
	}
	return b.String()
}

func main() {
	fmt.Println(maxSumOfSquares(2, 3))  // "30"
	fmt.Println(maxSumOfSquares(2, 17)) // "98"
	fmt.Println(maxSumOfSquares(1, 10)) // ""
}
