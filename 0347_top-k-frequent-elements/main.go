package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	topsByNum := make(map[int]int, k)
	topsByCount := make(map[int][]int, k)

	maxCount := 0

	for _, n := range nums {
		topsByNum[n] += 1

		if maxCount < topsByNum[n] {
			maxCount = topsByNum[n]
		}

		sameCount := topsByCount[topsByNum[n]]
		if sameCount == nil {
			sameCount = make([]int, 0, 1)
		}

		topsByCount[topsByNum[n]] = append(sameCount, n)
	}

	tops := make(map[int]struct{}, k)

	for i := maxCount; i > 0 && len(tops) < k; i-- {
		vals, ok := topsByCount[i]
		if !ok {
			continue
		}

		for _, v := range vals {
			tops[v] = struct{}{}
			if len(tops) == k {
				break
			}
		}

	}

	// fmt.Println(topsByNum)
	// fmt.Println(topsByCount)
	// fmt.Println(tops)

	res := make([]int, 0, len(tops))
	for k := range tops {
		res = append(res, k)
	}

	return res
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
