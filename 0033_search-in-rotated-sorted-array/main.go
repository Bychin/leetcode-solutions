package main

import (
	"fmt"
)

// 1) бинарным поиском ищем минимальный элемент (чтобы понять смещение)
// 2) бинарным поиском через фейковые индексы и смещение ищем нужный

func findRealStart(nums []int) int {
	min := 0
	max := len(nums) - 1

	pivot := (max + min) / 2

	for min < max {
		pivot = (max + min) / 2

		if pivot > 0 {
			if nums[pivot-1] > nums[pivot] {
				return pivot
			}
		} else {
			if nums[len(nums)-1] > nums[pivot] {
				return pivot
			}
		}

		if pivot < len(nums)-1 {
			if nums[pivot] > nums[pivot+1] {
				return pivot + 1
			}
		} else {
			if nums[pivot] > nums[0] {
				return 0
			}
		}

		if nums[pivot] > nums[len(nums)-1] {
			min = pivot
			continue
		}

		max = pivot
	}

	return pivot
}

func realIndex(index, realMin, len int) int {
	return (index + realMin) % len
}

func search(nums []int, target int) int {
	realMin := findRealStart(nums)

	min := 0             //realIndex(0, realMin, len(nums))
	max := len(nums) - 1 //realIndex(len(nums)-1, realMin, len(nums))

	for min <= max {
		pivot := (max + min) / 2
		value := nums[realIndex(pivot, realMin, len(nums))]

		if value == target {
			return realIndex(pivot, realMin, len(nums))
		}
		if value < target {
			min = pivot + 1
		}
		if value > target {
			max = pivot - 1
		}
	}

	return -1
}

func main() {
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
}
