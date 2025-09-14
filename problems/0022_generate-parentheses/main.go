package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	result := make([]string, 0, n)

	pattern := make([]string, 0)
	var fn func(openN, closeN int)
	fn = func(openN, closeN int) {
		if openN == n && closeN == n {
			result = append(result, strings.Join(pattern, ""))
			return
		}
		if openN < n {
			pattern = append(pattern, "(")
			fn(openN+1, closeN)
			pattern = pattern[:len(pattern)-1]
		}
		if closeN < openN {
			pattern = append(pattern, ")")
			fn(openN, closeN+1)
			pattern = pattern[:len(pattern)-1]
		}
	}
	fn(0, 0)

	return result
}

func main() {
	fmt.Println(generateParenthesis(3)) // ["((()))","(()())","(())()","()(())","()()()"]
	fmt.Println(generateParenthesis(1)) // ["()"]
}
