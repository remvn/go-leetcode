package main

import "fmt"

func main() {
	tests := [][]int{
		{2, 3, 1, 1, 4},
		{1, 3, 2},
		{1, 2, 1, 1, 1},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println("step:", jump(test))
		fmt.Println()
	}
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	index := 0
	step := 1

	// calculate the distance a nums[i] can jump
	// {10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0},
	// for example, the last 1 actually jump longer
	// than 9
	distance := func(index int, i int) int {
		return i - index + nums[i]
	}

	for index+nums[index]+1 < len(nums) {
		maxLength := min(len(nums), index+nums[index]+1)
		jumpIndex := index + 1
		for i := index + 1; i < maxLength; i++ {
			if distance(index, jumpIndex) <= distance(index, i) {
				jumpIndex = i
			}
		}
		// fmt.Printf("index: %d maxIndex: %d jumpIndex: %d\n", index, maxIndex, jumpIndex)
		index = jumpIndex
		step++
	}

	return step
}
