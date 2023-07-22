package main

import (
	"fmt"
)

func maximumWealth(accounts [][]int) (result int) {
	for i := 0; i < len(accounts); i++ {
		currMax := 0
		for j := 0; j < len(accounts[i]); j++ {
			currMax += accounts[i][j]
		}

		if currMax > result {
			result = currMax
		}
	}

	return
}

func main() {
	fmt.Println(maximumWealth([][]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}))
}
