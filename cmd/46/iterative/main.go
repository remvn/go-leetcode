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

type State struct {
	arr  []int
	nums []int
}

// iterative approach

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	numLength := len(nums)

	stack := []State{
		{arr: []int{}, nums: nums},
	}

	for len(stack) > 0 {
		state := stack[0]
		stack = stack[1:]
		if len(state.arr) == numLength {
			result = append(result, state.arr)
			continue
		}
		for i := 0; i < len(state.nums); i++ {
			remainNums := make([]int, len(state.nums))
			copy(remainNums, state.nums)
			remainNums = slices.Delete(remainNums, i, i+1)

			stack = append(stack, State{
				arr:  cloneAndAppend(state.arr, state.nums[i]),
				nums: remainNums,
			})
		}
	}
	return result
}

func cloneAndAppend(arr []int, val int) []int {
	clone := make([]int, len(arr), len(arr)+1)
	copy(clone, arr)
	clone = append(clone, val)
	return clone
}
