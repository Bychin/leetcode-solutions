package main

import (
	"fmt"
)

func reverseWords(s string) string {
	b := []byte(s)

	i := 0
	for j := 1; j < len(s)+1; {
		if j < len(s) && s[j] != ' ' {
			j++
			continue
		}

		// swap symbols
		k := i
		l := j - 1
		for k < l {
			b[k], b[l] = b[l], b[k]
			k++
			l--
		}

		i = j + 1
		j = i + 1
	}

	return string(b)
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest")) // "s'teL ekat edoCteeL tsetnoc"
	fmt.Println(reverseWords("Mr Ding"))                     // "rM gniD"
}
