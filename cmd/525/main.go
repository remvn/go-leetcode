package main

import (
	"fmt"
)

func main() {
	tests := [][]int{
		{1, 1, 0, 1, 0, 1, 0},
		{0, 1, 0},
		{0, 1, 1, 0, 1, 1, 1, 0},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(findMaxLength(test))
		fmt.Println()
	}
}

func findMaxLength(nums []int) int {
	sumMap := map[int]int{}

	// num = 0 -> sum--
	// num = 1 -> sum++

	// 1 1 0 1 0 1 0
	// 1 2 1 2 1 2 1

	// 0  1  1  0  1  1  1  0
	//-1  0  1  0

	sum := 0
	longest := 0
	for i, num := range nums {
		if num == 0 {
			sum--
		} else {
			sum++
		}
		if sum == 0 {
			longest = max(longest, i+1)
		}
		if index, ok := sumMap[sum]; ok {
			length := i - index
			longest = max(length, longest)
		} else {
			sumMap[sum] = i
		}
	}

	return longest
}
