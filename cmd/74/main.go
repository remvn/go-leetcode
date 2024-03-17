package main

import "fmt"

func main() {
	tests := []TestCase{
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
		},
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 11,
		},
		{
			matrix: [][]int{{1}},
			target: 1,
		},
		{
			matrix: [][]int{{1, 1}},
			target: 0,
		},
		{
			matrix: [][]int{{1}, {3}},
			target: 3,
		},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(searchMatrix(test.matrix, test.target))
		fmt.Println()
	}
}

type TestCase struct {
	matrix [][]int
	target int
}

func searchMatrix(matrix [][]int, target int) bool {
	top := 0
	bottom := len(matrix) - 1
	for top <= bottom {
		mid := top + (bottom-top)/2
		num := matrix[mid][0]
		if num == target {
			return true
		}
		if num > target {
			bottom = mid - 1
		} else {
			top = mid + 1
		}
	}

	// 1 3 5 10 23
	top = max(0, top-1)

	left := 0
	right := len(matrix[0]) - 1
	for left <= right {
		mid := left + (right-left)/2
		num := matrix[top][mid]
		if num == target {
			return true
		}
		if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
