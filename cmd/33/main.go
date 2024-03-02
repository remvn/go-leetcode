package main

import "fmt"

func main() {
	tests := []TestCase{
		{
			arr:    []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
		},
		{
			arr:    []int{4, 5, 6, 7, 8, 1, 2, 3},
			target: 8,
		},
		{
			arr:    []int{3, 1},
			target: 1,
		},
		{
			arr:    []int{4, 1, 2, 3},
			target: 1,
		},
	}
	for _, test := range tests {
		fmt.Printf("%+v: \n", test)
		fmt.Printf("result: %d\n", search(test.arr, test.target))
		fmt.Println()
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
		fmt.Printf("Pivot: %d \n", pivot)
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
		} else if target <= nums[mid] {
			right = right - 1
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

// idea: search for index of smallest number
// and substract 1
func searchForPivotPoint(nums []int) int {
	// this array isn't rotated
	if nums[0] < nums[len(nums)-1] {
		return -1
	}

	left := 0
	right := len(nums) - 1

	//  4 5 6 7 8 1 2 3
	//  4 5 6 7 0 1 2
	//  4 5 0 1 2
	//  3 1
	for left < right {
		fmt.Println(nums[left : right+1])
		mid := left + (right-left)/2
		if nums[mid] > nums[len(nums)-1] {
			// keep pushing left to find
			// the smallest number
			left = mid + 1
		}
		if nums[mid] < nums[len(nums)-1] {
			right = mid
			// - If we use "right = mid - 1"
			//	 we will miss the smallest number in some case
			//   because [left : mid - 1] doesn't contain
			//   smallest number when nums[mid] == smallest
			//   eg: [4,1,2,3] -> [4].
			// - If we assign right = nums[mid] == smallest, it
			//   will stay fixed until the end of the loop.
			//   So we can leave it there and pushing left
			//   until left == right
		}
	}

	return max(0, left-1)
}
