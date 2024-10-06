package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 1, 2}
	minNum := findMin(nums)
	fmt.Println(minNum)
}

// 1 2 3 4
// 4 1 2 3
// 3 4 1 2
// 2 3 4 1

// 1 2 3
// 3 1 2
// 2 3 1

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	for len(nums) > 2 {
		left := nums[0]
		right := nums[len(nums)-1]
		mid := nums[len(nums)/2]
		if mid > left && mid > right {
			nums = nums[len(nums)/2:]
		}
		if mid < left && mid < right {
			nums = nums[:len(nums)/2+1]
		}

		// array is sorted
		if mid > left && mid < right {
			return nums[0]
		}
	}

	if nums[0] > nums[1] {
		return nums[1]
	} else {
		return nums[0]
	}
}
