package main

import "fmt"

func main() {
	tests := []TestCase{
		{
			arr:    []int{5, 7, 7, 8, 8, 10},
			target: 8,
		},
		{
			arr:    []int{5, 7, 7, 8, 8, 10},
			target: 6,
		},
		{
			arr:    []int{1},
			target: 0,
		},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(searchRange(test.arr, test.target))
	}
}

type TestCase struct {
	arr    []int
	target int
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	low := searchLowBound(nums, target)
	high := searchHighBound(nums, target)

	return []int{low, high}
}

func searchLowBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else if target == nums[mid] {
			right = mid - 1
			// keep pushing right until nums[right + 1] == target
			// eg: target = 2
			// 1 2 2 2 3 3 3
			// |
			// right, left
			// so eventually left will end up in front of
			// right and the loop is finish
			// that's why we need to check left <= right
			// instead of left < right in normal binary search
		}
	}

	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func searchHighBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else if target == nums[mid] {
			// keep pushing left until nums[left - 1] == target
			// eg: target = 2
			// ... 2 2 2 3 3 3
			//           |   |
			//          left right
			// so eventually right will end up just right
			// behind left and the loop is finish
			// that's why we need to check left <= right
			// instead of left < right in normal binary search
			left = mid + 1
		}
	}

	if right == -1 || nums[right] != target {
		return -1
	}
	return right
}
