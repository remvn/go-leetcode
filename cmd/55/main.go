package main

import "fmt"

func main() {
	tests := [][]int{
		{2, 3, 1, 1, 4},
		{3, 2, 1, 0, 4},
		{2, 0, 0},
		{1, 3, 3, 5, 3, 0, 0, 4, 0, 0},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(canJump(test))
		fmt.Println()
	}
}

// basically find the max distance we could reach
// at current index
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	index := 0
	distance := func(index int, jumpIndex int) int {
		return jumpIndex - index + nums[jumpIndex]
	}

	// 1 3 3 5 3 0 0 0 0 0
	for index+nums[index] < len(nums)-1 {
		if nums[index] == 0 {
			return false
		}
		maxDistance := 0
		jumpIndex := index + 1
		minLength := min(len(nums)-1, index+nums[index]+1)
		for i := index + 1; i < minLength; i++ {
			jumpDistance := distance(index, i)
			if jumpDistance >= maxDistance {
				maxDistance = jumpDistance
				jumpIndex = i
			}
		}
		index = jumpIndex
	}

	return true
}
