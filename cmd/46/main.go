package main

import (
	"fmt"
	"slices"
)

func main() {
	tests := [][]int{
		{1, 2, 3},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(permute(test))
		fmt.Println()
	}
}

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	backtrack([]int{}, len(nums), nums, &result)
	return result
}

func backtrack(arr []int, length int, nums []int, result *[][]int) {
	if len(arr) == length {
		*result = append(*result, arr)
		return
	}

	// fmt.Printf("arr: %v, nums: %v \n", arr, nums)

	for i := 0; i < len(nums); i++ {
		remainNums := make([]int, len(nums))
		copy(remainNums, nums)
		remainNums = slices.Delete(remainNums, i, i+1)
		backtrack(cloneAndAppend(arr, nums[i]), length, remainNums, result)
	}
}

func cloneAndAppend(arr []int, val int) []int {
	clone := make([]int, len(arr), len(arr)+1)
	copy(clone, arr)
	clone = append(clone, val)
	return clone
}
