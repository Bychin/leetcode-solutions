package main

import (
	"fmt"
	"maps"
)

func checkInclusion(s1, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	s1Counts := map[byte]int{}
	s2Counts := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		s1Counts[s1[i]]++
		s2Counts[s2[i]]++
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		if maps.Equal(s1Counts, s2Counts) {
			return true
		}
		// shift s2Counts
		s2Counts[s2[i]]--
		if s2Counts[s2[i]] == 0 {
			delete(s2Counts, s2[i])
		}
		s2Counts[s2[i+len(s1)]]++
	}
	return maps.Equal(s1Counts, s2Counts)
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo")) // true
	fmt.Println(checkInclusion("ab", "eidboaoo")) // false
	fmt.Println(checkInclusion("adc", "dcda"))    // true
}
