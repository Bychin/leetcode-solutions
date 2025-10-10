package main

import (
	"fmt"
)

func isAlpha(char byte) bool {
	switch {
	case char >= 'a' && char <= 'z':
		return true
	case char >= 'A' && char <= 'Z':
		return true
	default:
		return false
	}
}

func reverseOnlyLetters(s string) string {
	b := []byte(s)

	i, j := 0, len(b)-1

	for i < j {
		if !isAlpha(b[i]) {
			i++
			continue
		}
		if !isAlpha(b[j]) {
			j--
			continue
		}

		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	return string(b)
}

func main() {
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))        // "j-Ih-gfE-dCba"
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!")) // "Qedo1ct-eeLg=ntse-T!"
}
