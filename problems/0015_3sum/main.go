package main

import (
	"fmt"
	"slices"
)

// sort
// x + y + z = 0 -> x + y = -z
// so for every element use it as z and find x & y using 2sum solution

func twoSumWithNegation(nums []int, targetIndex int) map[[2]int]struct{} {
	i, j := 0, len(nums)-1
	target := -nums[targetIndex]
	result := map[[2]int]struct{}{}

	for i < j {
		if i == targetIndex {
			i++
			continue
		}
		if j == targetIndex {
			j--
			continue
		}
		sum := nums[i] + nums[j]
		if sum == target {
			key := [2]int{i, j}
			if nums[key[0]] > nums[key[1]] {
				key[0], key[1] = key[1], key[0]
			}
			result[key] = struct{}{}
			i++
			j--
			continue
		}
		if sum > target {
			j--
			continue
		}
		if sum < target {
			i++
			continue
		}
	}

	return result
}

func threeSum(nums []int) (results [][]int) {
	slices.Sort(nums)

	set := map[[3]int]struct{}{}
	for i, z := range nums {
		if i > 0 && z == nums[i-1] {
			continue
		}

		twoSumResults := twoSumWithNegation(nums, i)
		for resultPair := range twoSumResults {
			solution := [3]int{nums[resultPair[0]], nums[resultPair[1]], z}
			if z < nums[resultPair[0]] {
				solution[0], solution[2] = solution[2], solution[0]
				solution[1], solution[2] = solution[2], solution[1]
			} else if z < nums[resultPair[1]] {
				solution[1], solution[2] = solution[2], solution[1]
			}
			if _, ok := set[solution]; !ok {
				set[solution] = struct{}{}
				results = append(results, solution[:])
			}
		}
	}

	return results
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{0, 1, 1}))             // []
	fmt.Println(threeSum([]int{0, 0, 0}))             // [[0,0,0]]

	fmt.Println(threeSum([]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10}))                      // [[-10,5,5],[-5,0,5],[-4,2,2],[-3,-2,5],[-3,1,2],[-2,0,2]]
	fmt.Println(threeSum([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})) // [[0,0,0]]
}
