package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ") // I'm too lazy to do this manually on the fly
	if len(words) != len(pattern) {
		return false
	}
	dict := map[byte]string{}
	reverseDict := map[string]byte{}
	for i := 0; i < len(pattern); i++ {
		char := pattern[i]
		word := words[i]
		wordFromDict, ok := dict[char]
		if ok {
			if wordFromDict != word {
				return false
			}
		} else {
			dict[char] = word
		}
		charFromDict, ok := reverseDict[word]
		if ok {
			if charFromDict != char {
				return false
			}
		} else {
			reverseDict[word] = char
		}
	}
	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))  // true
	fmt.Println(wordPattern("abba", "dog dog dog dog"))  // false
	fmt.Println(wordPattern("abba", "dog cat cat fish")) // false
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))  // false
}
