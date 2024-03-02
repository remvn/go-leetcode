package main

import "fmt"

func main() {
	tests := []TestCase{
		{
			arr:    []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
		},
		{
			arr:    []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
		},
		{
			arr:    []int{3, 1},
			target: 1,
		},
	}
	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(search(test.arr, test.target))
	}
}

type TestCase struct {
	arr    []int
	target int
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	pivot := searchForPivotPoint(nums)

	if pivot >= 0 {
		// check if target is inside pivot point
		if target >= nums[0] && target <= nums[pivot] {
			right = pivot
		} else {
			left = pivot + 1
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else if target == nums[mid] {
			left = mid
			break
		}
	}

	if left < len(nums) && nums[left] == target {
		return left
	} else {
		return -1
	}
}

func searchForPivotPoint(nums []int) int {
	// this array isn't rotated
	if nums[0] < nums[len(nums)-1] {
		return -1
	}

	left := 0
	right := len(nums) - 1
	// 4,5,6,7,0,1,2
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[len(nums)-1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return max(0, left-1)
}
