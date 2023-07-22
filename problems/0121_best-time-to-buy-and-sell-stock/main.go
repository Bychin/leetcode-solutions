package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice, maxProfit := 0, 0

	for i, price := range prices {
		if i == 0 {
			minPrice = prices[i]
			continue
		}

		if price < minPrice {
			minPrice = price
			continue
		}

		if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
