package main

import "fmt"

func main() {
	tests := []TestCase{
		{
			arr:    []int{1, 3, 5, 6},
			target: 4,
		},
		{
			arr:    []int{1, 3, 5, 6},
			target: 7,
		},
		{
			arr:    []int{1, 3, 5, 6},
			target: 0,
		},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(searchInsert(test.arr, test.target))
		fmt.Println()
	}
}

type TestCase struct {
	arr    []int
	target int
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	// [1,3,5,6]
	// 5
	// [1,3,5,6]
	// 4
	// [1,3,5,6]
	// 7

	// the trick here is to use
	// for left <= right
	// to make left = len(num)
	// if we insert number at the end
	// of array
	for left <= right {
		mid := left + (right-left)/2
		// fmt.Println(nums[left : right+1])
		if target > nums[mid] {
			left = mid + 1
		}
		if target < nums[mid] {
			right = mid - 1
		}
		if target == nums[mid] {
			return mid
		}
	}

	return left
}
