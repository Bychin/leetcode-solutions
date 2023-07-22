package main

import (
	"fmt"
)

// 1) ищем стартующий элемент (обычным поиском, тк здесь возможно дублирование)
// 2) бинарным поиском через фейковые индексы и смещение ищем нужный

func findRealStart(nums []int) int {
	pivot := len(nums) - 1
	minValue := 0

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			minValue = nums[i]
			continue
		}

		if nums[i] > minValue {
			return i + 1
		}
		if nums[i] <= minValue {
			minValue = nums[i]
			pivot = i
		}
	}

	return pivot
}

func realIndex(index, realMin, len int) int {
	return (index + realMin) % len
}

func search(nums []int, target int) bool {
	realMin := findRealStart(nums)

	fmt.Println(realMin)

	min := 0             //realIndex(0, realMin, len(nums))
	max := len(nums) - 1 //realIndex(len(nums)-1, realMin, len(nums))

	for min <= max {
		pivot := (max + min) / 2
		value := nums[realIndex(pivot, realMin, len(nums))]

		if value == target {
			return true
		}
		if value < target {
			min = pivot + 1
		}
		if value > target {
			max = pivot - 1
		}
	}

	return false
}

func main() {
	fmt.Println(search([]int{1, 1, 1, 2, 1, 1}, 2))
	fmt.Println(search([]int{3, 5, 6, 0, 1, 2}, 0))
	fmt.Println(search([]int{3, 5, 6, 0, 1, 2}, 3))
	fmt.Println(search([]int{3, 5, 6, 0, 1, 2}, 2))
	fmt.Println(search([]int{1, 2, 3, 4, 5, 6}, 0))
	fmt.Println(search([]int{6, 1, 2, 3, 4, 5}, 0))
	fmt.Println(search([]int{3, 4, 5, 6, 1, 2}, 0))
	fmt.Println(search([]int{2, 3, 4, 5, 6, 1}, 0))
	fmt.Println(search([]int{1}, 1))
	fmt.Println(search([]int{1}, 2))
	fmt.Println(search([]int{1, 2}, 2))
	fmt.Println(search([]int{1, 2, 3}, 2))
	fmt.Println(search([]int{1, 1, 1, 1, 3}, 2))
	fmt.Println(search([]int{1, 1, 1, 1, 3}, 3))
}
