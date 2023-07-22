package main

import (
	"fmt"
)

type point struct {
	i, j int
}

func dfs(grid [][]byte, visited map[point]struct{}, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}

	if grid[i][j] == '0' {
		return
	}

	if _, ok := visited[point{i, j}]; ok {
		return
	}

	visited[point{i, j}] = struct{}{}

	dfs(grid, visited, i+1, j)
	dfs(grid, visited, i-1, j)
	dfs(grid, visited, i, j+1)
	dfs(grid, visited, i, j-1)
}

func numIslands(grid [][]byte) (result int) {
	if len(grid) == 0 {
		return 0
	}

	visited := make(map[point]struct{}, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				if _, ok := visited[point{i, j}]; ok {
					continue
				}

				result += 1

				dfs(grid, visited, i, j)
			}
		}
	}

	return
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))

	fmt.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
}
