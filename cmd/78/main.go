package main

import "fmt"

func main() {
	tests := [][]int{
		{1, 2, 3}, {0},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(subsets(test))
		fmt.Println()
	}
}

func subsets(nums []int) [][]int {
	result := [][]int{}
	backtracking([]int{}, nums, &result)
	return result
}

func backtracking(arr []int, nums []int, result *[][]int) {
	*result = append(*result, arr)
	if len(nums) == 0 {
		return
	}

	for i, num := range nums {
		clone := cloneAndAppend(arr, num)
		backtracking(clone, nums[i+1:], result)
	}
}

func cloneAndAppend(arr []int, val int) []int {
	clone := make([]int, len(arr))
	copy(clone, arr)
	clone = append(clone, val)
	return clone
}
