package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := []TestCase{
		// {
		// 	arr:    []int{2, 3, 6, 7},
		// 	target: 7,
		// },
		{
			arr:    []int{2, 3, 5},
			target: 8,
		},
	}

	for _, test := range tests {
		fmt.Printf("%+v \n", test)
		fmt.Println(combinationSum(test.arr, test.target))
		fmt.Println()
	}
}

type TestCase struct {
	arr    []int
	target int
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	combination([]int{}, candidates, target, &result)
	return result
}

func combination(arr []int, candidates []int, target int, result *[][]int) {
	// fmt.Println(arr)
	total := sumOfArr(arr)
	if total == target {
		// fmt.Println("append: ", arr)
		*result = append(*result, arr)
		return
	}
	if total > target {
		return
	}

	if len(arr) > 0 {
		// repeat last digit until a + a + a + ... + 1 > total
		// then the loop below will append new digit (eg: a + a + a + b)
		// and start this loop again
		combination(cloneAndAppend(arr, arr[len(arr)-1]), candidates, target, result)
	}

	for i := 0; i < len(candidates); i++ {
		curr := candidates[i]
		if target-curr < 0 {
			break
		}
		combination(cloneAndAppend(arr, curr), candidates[i+1:], target, result)
	}
}

func cloneAndAppend(arr []int, num int) []int {
	clone := make([]int, len(arr), len(arr)+1)
	copy(clone, arr)
	clone = append(clone, num)
	return clone
}

func sumOfArr(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}
