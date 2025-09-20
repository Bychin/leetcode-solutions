package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum < target {
			i++
			continue
		}
		if sum > target {
			j--
			continue
		}
	}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // {1, 2}
	fmt.Println(twoSum([]int{2, 3, 4}, 6))      // {1, 3}
	fmt.Println(twoSum([]int{-1, 0}, -1))       // {1, 2}
}
