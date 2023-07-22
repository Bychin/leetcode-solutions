package main

import (
	"fmt"
)

func subtractProductAndSum(n int) int {
	product := 1
	sum := 0

	for n > 0 {
		k := n % 10
		n /= 10

		sum += k
		product *= k
	}

	return product - sum
}

func main() {
	fmt.Println(subtractProductAndSum(4421))
}
