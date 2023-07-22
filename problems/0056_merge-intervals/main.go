package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	result := make([][]int, 0, len(intervals))

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	currRange := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if currRange[1] >= intervals[i][0] {
			currRange[1] = max(currRange[1], intervals[i][1])
		} else {
			result = append(result, currRange)
			currRange = intervals[i]
		}
	}

	result = append(result, currRange)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(merge([][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}))
	fmt.Println(merge([][]int{
		{1, 4},
		{5, 6},
	}))
	fmt.Println(merge([][]int{
		{1, 4},
		{0, 4},
	}))
	fmt.Println(merge([][]int{
		{1, 4},
		{2, 3},
	}))
	fmt.Println(merge([][]int{
		{1, 4},
		{0, 2},
		{3, 5},
	}))
}
