package main

import "fmt"

func scalarIndexToVector(index, colsN int) (int, int) {
	i := index / colsN
	j := index % colsN

	return i, j
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}

	scalarMin := 0
	scalarMax := len(matrix) * len(matrix[0])

	for scalarMin <= scalarMax {
		pivot := (scalarMax + scalarMin) / 2
		i, j := scalarIndexToVector(pivot, len(matrix[0]))
		if i > len(matrix)-1 || j > len(matrix[0])-1 {
			return false
		}
		value := matrix[i][j]

		if value == target {
			return true
		}
		if value < target {
			scalarMin = pivot + 1
		}
		if value > target {
			scalarMax = pivot - 1
		}
	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 16))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	fmt.Println(searchMatrix([][]int{{1}}, 2))
}
