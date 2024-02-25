package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{0, 1, 1},
		{0, 0, 0},
		{-2, 0, 0, 2, 2},
	}
	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(threeSum(test))
	}
}

func threeSum(nums []int) [][]int {
	var result [][]int

	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			target := nums[left] + nums[right] + nums[i]
			if target == 0 {
				result = append(result, []int{nums[left], nums[right], nums[i]})
				left++
				right--

				// check if next element is the same
				// the skip it
				for left < right && nums[left-1] == nums[left] {
					left++
				}
				for right > left && nums[right+1] == nums[right] {
					right--
				}
			} else {
				if target > 0 {
					right--
				} else {
					left++
				}
			}
		}
	}

	return result
}
