package main

import (
	"fmt"
	"slices"
)

type State struct {
	arr  []int
	nums []int
}

func main() {
	inputNums := []int{0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	numLength := 2

	nums := []int{}
	for i := 0; i < len(inputNums); i++ {
		if i > 0 && inputNums[i-1] == inputNums[i] {
			continue
		}
		nums = append(nums, inputNums[i])
	}

	result := [][]int{}
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
			if len(state.arr) == 0 && state.nums[i] == 0 {
				continue
			}
			stack = append(stack, State{
				arr:  append(slices.Clone(state.arr), state.nums[i]),
				nums: nums,
			})
		}
	}

	for _, item := range result {
		fmt.Println(item)
	}

}
