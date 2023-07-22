package main

import (
	"fmt"
)

func average(salary []int) float64 {
	minSalary, maxSalary := 0, 0

	sum := 0

	for _, s := range salary {
		sum += s
		if s > maxSalary {
			maxSalary = s
		}
		if s < minSalary || minSalary == 0 {
			minSalary = s
		}
	}

	sum -= (minSalary + maxSalary)

	return float64(sum) / float64(len(salary)-2)
}

func main() {
	fmt.Println(average([]int{4000, 3000, 1000, 2000}))
	fmt.Println(average([]int{1000, 2000, 3000}))
	fmt.Println(average([]int{1000, 1000, 1000}))
	fmt.Println(average([]int{1000, 1000, 2000}))
	fmt.Println(average([]int{1000, 2000, 2000, 1000}))
}
