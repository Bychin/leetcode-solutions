package main

import (
	"fmt"
	"math"
)

func smallestAbsent(nums []int) int {
	sum := 0
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		sum += num
		set[num] = struct{}{}
	}
	average := float64(sum) / float64(len(nums))

	result := int(math.Ceil(average))
	if average == math.Ceil(average) {
		result = int(math.Ceil(average + 1))
	}
	if result < 1 {
		result = 1
	}

	for {
		if _, ok := set[result]; !ok {
			break
		} else {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(smallestAbsent([]int{0, 1, 2, 3})) // Output: 4
	fmt.Println(smallestAbsent([]int{-1, 1, 2}))   // Output: 3
	fmt.Println(smallestAbsent([]int{3, 5}))       // Output: 6
	fmt.Println(smallestAbsent([]int{-1, 1, 2}))   // Output: 3
	fmt.Println(smallestAbsent([]int{4, -1}))      // Output: 2
	fmt.Println(smallestAbsent([]int{-34}))        // Output: 1
	fmt.Println(smallestAbsent([]int{0}))          // Output: 1
	fmt.Println(smallestAbsent([]int{1, -39, 9}))  // Output: 2
}
