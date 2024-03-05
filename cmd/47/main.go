package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func main() {
	tests := [][]int{
		{1, 2, 3},
		{1, 1, 2},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(permuteUnique(test))
		fmt.Println()
	}
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	backtrack([]int{}, nums, len(nums), &result)

	return result
}

func backtrack(arr []int, nums []int, length int, result *[][]int) {
	// fmt.Printf("arr: %v %s remain: %v\n", arr, strings.Repeat(" ", 8-max(2*len(arr)-1, 0)), nums)
	if len(arr) == length {
		*result = append(*result, arr)
		return
	}

	// 1 1 2
	// 1
	//   1 2
	//   2 1
	// 2
	//   1 1
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		remainNums := make([]int, len(nums))
		copy(remainNums, nums)
		remainNums = slices.Delete(remainNums, i, i+1)
		backtrack(cloneAndAppend(arr, nums[i]), remainNums, length, result)
	}
}

func cloneAndAppend(arr []int, val int) []int {
	clone := make([]int, len(arr), len(arr)+1)
	copy(clone, arr)
	clone = append(clone, val)
	return clone
}
