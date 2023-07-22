package main

import (
	"fmt"
)

func hammingWeight(num uint32) (counter int) {
	// return bits.OnesCount32(num)

	for num > 0 {
		counter++
		num &= (num - 1)
	}

	return
}

func main() {
	fmt.Println(hammingWeight(4294967293))
}
