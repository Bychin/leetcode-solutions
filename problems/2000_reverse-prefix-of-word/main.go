package main

import (
	"fmt"
)

func reversePrefix(word string, ch byte) string {
	wordAsBytes := []byte(word)
	prefixEnd := -1
	for i := 0; i < len(wordAsBytes); i++ {
		if wordAsBytes[i] == ch {
			prefixEnd = i
			break
		}
	}
	prefixStart := 0
	for prefixStart < prefixEnd {
		wordAsBytes[prefixStart], wordAsBytes[prefixEnd] = wordAsBytes[prefixEnd], wordAsBytes[prefixStart]
		prefixStart++
		prefixEnd--
	}
	return string(wordAsBytes)
}

func main() {
	fmt.Println(reversePrefix("abcdefd", 'd')) // "dcbaefd"
	fmt.Println(reversePrefix("xyxzxe", 'z'))  // "zxyxxe"
	fmt.Println(reversePrefix("abcd", 'z'))    // "abcd"
}
