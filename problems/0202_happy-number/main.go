package main

import (
	"fmt"
)

func calculate(n int) int {
	result := 0

	for n != 0 {
		value := n % 10
		result += value * value
		n = n / 10
	}

	return result
}

func isHappy(n int) bool {
	slow := n
	fast := calculate(n)

	for slow != 1 && slow != fast {
		slow = calculate(slow)
		fast = calculate(calculate(fast))
	}

	return fast == 1
}

func main() {
	fmt.Println(isHappy(19)) // true
	fmt.Println(isHappy(2))  // false
}
