package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	// valuesInSubBox maps sub-box i index to j index to unique values in the sub-box
	valuesInSubBox := make(map[int]map[int]map[byte]struct{}, 3)
	for i := 0; i < 3; i++ {
		valuesInSubBox[i] = make(map[int]map[byte]struct{}, 3)
		for j := 0; j < 3; j++ {
			valuesInSubBox[i][j] = make(map[byte]struct{}, 9)
		}
	}

	for i := 0; i < len(board); i++ {
		valuesInRow := make(map[byte]struct{}, len(board))
		valuesInColumn := make(map[byte]struct{}, len(board))

		for j := 0; j < len(board); j++ {
			// rows
			digit := board[i][j]
			if _, ok := valuesInRow[digit]; ok && digit != '.' {
				return false
			}
			valuesInRow[digit] = struct{}{}

			// columns
			digit = board[j][i]
			if _, ok := valuesInColumn[digit]; ok && digit != '.' {
				return false
			}
			valuesInColumn[digit] = struct{}{}

			// sub-boxes
			subBoxJ := j / 3
			subBoxI := i / 3
			if _, ok := valuesInSubBox[subBoxJ][subBoxI][digit]; ok && digit != '.' {
				return false
			}
			valuesInSubBox[subBoxJ][subBoxI][digit] = struct{}{}
		}
	}

	return true
}

var testCases = [][][]byte{
	{ // true
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	},
	{ // false
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	},
}

func main() {
	for _, tc := range testCases {
		fmt.Println(isValidSudoku(tc))
	}
}
