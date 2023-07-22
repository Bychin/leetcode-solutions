package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	betterMagazine := map[rune]int{}

	for _, char := range magazine {
		counter := betterMagazine[char]
		counter++
		betterMagazine[char] = counter
	}

	for _, char := range ransomNote {
		counter := betterMagazine[char]
		if counter == 0 {
			return false
		}
		counter--
		betterMagazine[char] = counter
	}

	return true
}

func main() {
	fmt.Println(canConstruct("aa", "aab"))
}
