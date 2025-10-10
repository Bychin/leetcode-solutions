package main

import (
	"fmt"
)

func destCity(paths [][]string) string {
	pathsAsMap := make(map[string]string, len(paths))

	for _, p := range paths {
		start, end := p[0], p[1]
		pathsAsMap[start] = end
	}

	for _, p := range paths {
		end := p[1]
		if _, ok := pathsAsMap[end]; !ok {
			return end
		}
	}

	return ""
}

func main() {
	fmt.Println(destCity([][]string{
		{"London", "New York"},
		{"New York", "Lima"},
		{"Lima", "Sao Paulo"},
	})) // "Sao Paulo"
	fmt.Println(destCity([][]string{
		{"B", "C"},
		{"D", "B"},
		{"C", "A"},
	})) // "A"
	fmt.Println(destCity([][]string{
		{"A", "Z"},
	})) // "Z"
}
