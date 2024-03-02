package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := []TestCase{
		{arr: []int{10, 1, 2, 7, 6, 1, 5}, target: 8},
		{arr: []int{2, 5, 2, 1, 2}, target: 7},
	}
	for _, test := range tests {
		fmt.Printf("test: %+v \n", test)
		fmt.Println(combinationSum2(test.arr, test.target))
		fmt.Println()
	}
}

type TestCase struct {
	arr    []int
	target int
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	// fmt.Println("candidates: ", candidates)

	result := [][]int{}
	backtrack([]int{}, candidates, target, &result)

	return result
}

func backtrack(arr []int, candidates []int, target int, result *[][]int) {
	// fmt.Println(arr)
	total := sumOfArr(arr)
	if total == target {
		*result = append(*result, arr)
		return
	}
	if total > target {
		return
	}

	for i := 0; i < len(candidates); i++ {
		curr := candidates[i]
		if target-curr < 0 {
			break
		}
		// eg: 1 2 2 2 5
		// 1 - first loop
		// 1 2 - second loop
		// 1 2 2 - third loop
		// ...
		// 1 2 - second loop  * skip this with nums[i] == nums[i-1]
		// 1 2 - second loop  * skip this with nums[i] == nums[i-1]
		// 2 - first loop
		// ...
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		backtrack(cloneAndAppend(arr, curr), candidates[i+1:], target, result)
	}
}

func cloneAndAppend(arr []int, val int) []int {
	clone := make([]int, len(arr), len(arr)+1)
	copy(clone, arr)
	clone = append(clone, val)
	return clone
}

func sumOfArr(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}
