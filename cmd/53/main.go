package main

import (
	"fmt"
)

func main() {
	tests := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1},
		{5, 4, -1, 7, 8},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println("result:", maxSubArray(test))
		fmt.Println()
	}
}

func maxSubArray(nums []int) int {
	sum, result := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			// if this current sum is negative then total of
			// sum + sumOfNextSubArray (i, i+1, i+2)
			// is clearly smaller than sumOfNextSubArray.
			// In other word, this current sum is now useless
			// and we better start over (with sum = 0)
			sum = 0
		}
		sum += nums[i]
		if sum > result {
			result = sum
		}
		// fmt.Println(sum)
	}
	return result
}
