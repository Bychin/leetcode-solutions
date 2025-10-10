package main

import (
	"fmt"
)

func checkPoint(i, j int, visited map[[2]int]struct{}) bool {
	point := [2]int{i, j}
	if _, ok := visited[point]; ok {
		return true
	}
	visited[point] = struct{}{}
	return false
}

func isPathCrossing(path string) bool {
	i, j := 0, 0
	visited := map[[2]int]struct{}{
		{i, j}: {},
	}

	for _, p := range path {
		switch p {
		case 'N':
			i++
			if checkPoint(i, j, visited) {
				return true
			}
		case 'S':
			i--
			if checkPoint(i, j, visited) {
				return true
			}
		case 'E':
			j++
			if checkPoint(i, j, visited) {
				return true
			}
		case 'W':
			j--
			if checkPoint(i, j, visited) {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(isPathCrossing("NES"))   // false
	fmt.Println(isPathCrossing("NESWW")) // true
}
