package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := []TestCase{
		{arr: []int{10, 1, 2, 7, 6, 1, 5}, target: 8},
		{arr: []int{2, 5, 2, 1, 2}, target: 5},
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
	result := combination(candidates, target)
	return result
}

func combination(candidates []int, target int) [][]int {
	fmt.Println("candidates: ", candidates)
	result := [][]int{}
	for i := 0; i < len(candidates); i++ {
		fmt.Println("i:", i)
		recursiveAppend([]int{candidates[i]}, candidates[i+1:], target, &result)

		for i+1 < len(candidates) && candidates[i+1] == candidates[i] {
			i++
		}
	}

	return result
}

func recursiveAppend(arr []int, candidates []int, target int, result *[][]int) {
	fmt.Println(arr)
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
		recursiveAppend(cloneAndAppend(arr, curr), candidates[i+1:], target, result)
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
