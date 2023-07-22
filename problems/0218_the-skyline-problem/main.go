package main

import (
	"fmt"
	"sort"
)

func insertAt(data []int, i int, v int) []int {
	if i == len(data) {
		return append(data, v)
	}

	data = append(data[:i+1], data[i:]...)
	data[i] = v

	return data
}

func insertSorted(data []int, v int) []int {
	i := sort.Search(len(data), func(i int) bool { return data[i] >= v })
	return insertAt(data, i, v)
}

func removeFromSorted(data []int, v int) []int {
	i := sort.Search(len(data), func(i int) bool { return data[i] >= v })
	if i == len(data) {
		panic("unexpected remove")
	}

	if i == len(data)-1 {
		return data[:i]
	}

	return append(data[:i], data[i+1:]...)
}

type Point struct {
	x int
	y int // negative sign for going up
}

func getSkyline(buildings [][]int) [][]int {
	points := make([]Point, 0, 2*len(buildings))

	for _, b := range buildings {
		points = append(points,
			Point{
				x: b[0],
				y: -b[2],
			},
			Point{
				x: b[1],
				y: b[2],
			},
		)
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].x != points[j].x {
			return points[i].x < points[j].x
		}

		return points[i].y < points[j].y
	})

	// fmt.Println(points)

	skyline := make([][]int, 0)

	sortedHeights := make([]int, 0, len(points))
	currentHeight := 0

	for i, p := range points {
		if i == 0 {
			skyline = append(skyline, []int{p.x, -p.y})

			sortedHeights = insertSorted(sortedHeights, -p.y)
			currentHeight = -p.y

			continue
		}

		if p.y < 0 { // y always != 0
			sortedHeights = insertSorted(sortedHeights, -p.y)
		} else {
			sortedHeights = removeFromSorted(sortedHeights, p.y)
		}

		maxHeight := 0
		if len(sortedHeights) > 0 {
			maxHeight = sortedHeights[len(sortedHeights)-1]
		}
		if maxHeight != currentHeight {
			currentHeight = maxHeight
			skyline = append(skyline, []int{p.x, currentHeight})
		}

	}

	return skyline
}

func main() {
	fmt.Println(getSkyline([][]int{
		{2, 9, 10}, // x_left, x_right, y
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}))

	fmt.Println(getSkyline([][]int{
		{0, 2, 3}, // x_left, x_right, y
		{2, 5, 3},
	}))

	fmt.Println(getSkyline([][]int{
		{1, 2, 1},
		{1, 2, 2},
		{1, 2, 3},
	}))

	fmt.Println(getSkyline([][]int{
		{3, 7, 8}, // x_left, x_right, y
		{3, 8, 7},
		{3, 9, 6},
		{3, 10, 5},
		{3, 11, 4},
		{3, 12, 3},
		{3, 13, 2},
		{3, 14, 1},
	}))

}
