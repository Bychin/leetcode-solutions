package main

import (
	"fmt"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func lexSmallest(s string) string {
	smallest := min(s, reverse(s))

	for k := 1; k < len(s); k++ {
		smallest = min(smallest, s[:k]+reverse(s[k:]))
		smallest = min(smallest, reverse(s[:k])+s[k:])
	}

	return smallest
}

func main() {
	fmt.Println(lexSmallest("dcab")) // "acdb"
	fmt.Println(lexSmallest("abba")) // "aabb"
	fmt.Println(lexSmallest("zxy"))  // "xzy"
	fmt.Println(lexSmallest("caa"))  // "aac"
}
