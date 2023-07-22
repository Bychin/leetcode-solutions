package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i < n+1; i++ {
		if i%3 == 0 {
			result[i-1] += "Fizz"
		}
		if i%5 == 0 {
			result[i-1] += "Buzz"
		}
		if result[i-1] == "" {
			result[i-1] = strconv.Itoa(i)
		}
	}

	return result
}

func main() {
	fmt.Println(fizzBuzz(15))
}
