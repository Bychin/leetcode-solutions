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

		topsByCount[topsByNum[n]] = append(topsByCount[topsByNum[n]], n)
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

func topKFrequentUsingBucketSort(nums []int, k int) []int {
	counts := map[int]int{}
	maxCount := 0 // it is always <= len(nums)

	for _, n := range nums {
		counts[n] += 1
		maxCount = max(maxCount, counts[n])
	}

	buckets := make([][]int, maxCount+1)
	for n, count := range counts {
		buckets[count] = append(buckets[count], n)
	}

	result := []int{}
	for i := maxCount; i > 0 && k > 0; i-- {
		ns := buckets[i]
		for j := 0; j < len(ns) && k > 0; j++ {
			result = append(result, ns[j])
			k--
		}
	}

	return result
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))             // [1,2]
	fmt.Println(topKFrequent([]int{1}, 1))                            // [1]
	fmt.Println(topKFrequent([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2)) // [1,2]

	fmt.Println(topKFrequentUsingBucketSort([]int{1, 1, 1, 2, 2, 3}, 2))             // [1,2]
	fmt.Println(topKFrequentUsingBucketSort([]int{1}, 1))                            // [1]
	fmt.Println(topKFrequentUsingBucketSort([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2)) // [1,2]
}
