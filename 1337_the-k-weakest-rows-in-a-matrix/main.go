package main

import (
	"fmt"
)

func kWeakestRows(mat [][]int, k int) []int {
	powerToIndex := make(map[int][]int, len(mat))

	for i := 0; i < len(mat); i++ {
		j := 0
		for ; j < len(mat[i]) && mat[i][j] == 1; j++ {
			// pass
		}
		powerToIndex[j] = append(powerToIndex[j], i)
	}

	result := make([]int, 0, k)

	for i := 0; i <= len(mat[0]) && len(result) < k; i++ {
		toAdd := powerToIndex[i]
		if len(toAdd) > k-len(result) {
			toAdd = toAdd[:k-len(result)]
		}
		result = append(result, toAdd...)
	}

	return result
}

func kWeakestRowsVerticalTraverse(mat [][]int, k int) (result []int) {
	result = make([]int, 0, k)
	resultLookup := make(map[int]struct{})

	for j := 0; j < len(mat[0]); j++ {
		for i := 0; i < len(mat); i++ {
			if _, ok := resultLookup[i]; ok {
				continue
			}
			if mat[i][j] == 0 {
				result = append(result, i)
				if len(result) == k {
					return
				}
				resultLookup[i] = struct{}{}
			}
		}
	}

	if len(result) < k {
		for i := 0; len(result) != k; i++ {
			if _, ok := resultLookup[i]; !ok {
				result = append(result, i)
			}
		}
	}

	return
}

func main() {
	fmt.Println(kWeakestRows(
		[][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1},
		},
		3,
	))

	fmt.Println(kWeakestRows(
		[][]int{
			{1, 0},
			{0, 0},
			{1, 0},
		},
		2,
	))

	fmt.Println(kWeakestRowsVerticalTraverse(
		[][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1},
		},
		3,
	))

	fmt.Println(kWeakestRowsVerticalTraverse(
		[][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
		1,
	))
}
