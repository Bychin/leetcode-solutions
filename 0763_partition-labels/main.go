package main

import (
	"fmt"
)

func partitionLabels(s string) []int {
	firstAndLast := make(map[rune][2]int, 26)

	for i, char := range s {
		v, ok := firstAndLast[char]
		if !ok {
			firstAndLast[char] = [2]int{i, i}
			continue
		}

		if v[1] < i {
			firstAndLast[char] = [2]int{v[0], i}
		}
	}

	result := make([]int, 0)

	currPartitionStart := 0
	currPartitionEnd := 0
	for i, char := range s {
		v := firstAndLast[char]
		first, last := v[0], v[1]
		if i == 0 {
			currPartitionEnd = last
			continue
		}

		if first <= currPartitionEnd {
			if last > currPartitionEnd {
				currPartitionEnd = last
			}
			continue
		}

		result = append(result, currPartitionEnd-currPartitionStart+1)
		currPartitionStart = first
		currPartitionEnd = last
	}

	result = append(result, currPartitionEnd-currPartitionStart+1)
	return result
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
	fmt.Println(partitionLabels("vhaagbqkaq"))
}
