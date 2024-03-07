package main

import (
	"fmt"
	"slices"
)

func main() {
	tests := []TestCase{
		{
			intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newItem:   []int{4, 8},
		},
		{
			intervals: [][]int{{6, 8}},
			newItem:   []int{1, 5},
		},
		{
			intervals: [][]int{{2, 3}, {5, 5}, {6, 6}, {7, 7}, {8, 11}},
			newItem:   []int{6, 13},
		},
		{
			intervals: [][]int{{2, 6}, {7, 9}},
			newItem:   []int{15, 18},
		},
	}

	for _, test := range tests {
		fmt.Printf("case: %+v\n", test)
		fmt.Println("result:", insert(test.intervals, test.newItem))
		fmt.Println()
	}
}

type TestCase struct {
	intervals [][]int
	newItem   []int
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}

	left := 0
	right := len(intervals) - 1
	for left <= right {
		mid := left + (right-left)/2
		// fmt.Println(left, mid, right)
		// fmt.Println(intervals[left : right+1])
		// fmt.Println()
		if newInterval[0] >= intervals[mid][0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// fmt.Println("left: ", left)
	intervals = slices.Insert(intervals, left, newInterval)

	results := make([][]int, 0, len(intervals))
	for i := 0; i < len(intervals); i++ {
		if i > 0 && isConflict(results[len(results)-1], intervals[i]) {
			results[len(results)-1] = mergeArr(results[len(results)-1], intervals[i])
		} else {
			results = append(results, intervals[i])
		}
	}

	return results
}

func isConflict(a, b []int) bool {
	return a[1] >= b[0]
}

func mergeArr(a, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}
