package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	min := 1
	max := n

	for min <= max {
		pivot := (max + min) / 2
		r := guess(pivot)

		if r == 0 {
			return pivot
		}

		if r == -1 {
			max = pivot - 1
		}
		if r == 1 {
			min = pivot + 1
		}
	}

	return -1
}

func main() {
}
