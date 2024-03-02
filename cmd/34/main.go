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
		// {
		// 	arr:    []int{},
		// 	target: 0,
		// },
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

	start := 0
	end := len(nums) - 1

	// arr:    []int{5, 7, 7, 8, 8, 10},
	for start < end {
		if nums[start] == nums[end] && nums[start] == target {
			break
		}
		mid := (start + end) / 2
		fmt.Printf("%d %d %d\n", start, end, mid)
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			fmt.Printf("mid: %d", nums[mid])
			break
		}
	}

	if nums[start] == nums[end] && nums[start] == target {
		return []int{start, end}
	}
	return []int{-1, -1}
}
