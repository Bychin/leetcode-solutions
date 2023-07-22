package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	prefix := make([]byte, 0)

	for i := 0; ; i++ {
		counter := 0

		for j := 0; j < len(strs); j++ {
			if len(strs[j]) <= i {
				return string(prefix)
			}

			if strs[j][i] != strs[0][i] {
				return string(prefix)
			}

			counter++
		}

		if counter == len(strs) {
			prefix = append(prefix, strs[0][i])
		}
	}
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
