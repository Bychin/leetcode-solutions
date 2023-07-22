package main

import (
	"fmt"
)

var symbolsMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) (value int) {
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			value += symbolsMap[s[i]]
			return
		}

		if symbolsMap[s[i]] < symbolsMap[s[i+1]] {
			value -= symbolsMap[s[i]]
			continue
		}

		value += symbolsMap[s[i]]
	}

	return
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
