package main

import (
	"fmt"
)

type pair struct {
	t int
	i int
}

func dailyTemperatures(temperatures []int) []int {
	stack := []pair{}
	result := make([]int, len(temperatures))

	for i, t := range temperatures {
		for len(stack) > 0 && t > stack[len(stack)-1].t {
			last := stack[len(stack)-1]
			result[last.i] = i - last.i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{
			t: t,
			i: i,
		})
	}

	return result
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})) // [1,1,4,2,1,1,0,0]
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))                 // [1,1,1,0]
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))                     // [1,1,0]
}
