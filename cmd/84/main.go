package main

import (
	"fmt"
)

func main() {
	tests := [][]int{
		{2, 1, 5, 6, 2, 3},
		{2, 4},
		{5, 5, 1, 7, 1, 1, 5, 2, 7, 6},
		{8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783, 8783},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(largestRectangleArea(test))
		fmt.Println()
	}
}

func largestRectangleArea(heights []int) int {
	stack := [][]int{heights}
	largest := 0

	for len(stack) > 0 {
		fmt.Println(stack)
		arr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		area, minIndex, maxIndex := calculateArea(arr)
		largest = max(area, largest)

		if maxIndex == minIndex && maxIndex == len(arr)-1 {
			continue
		}

		arr1 := arr[minIndex+1:]
		arr2 := arr[:minIndex]
		if len(arr1) > 0 {
			stack = append(stack, arr1)
		}
		if len(arr2) > 0 {
			stack = append(stack, arr2)
		}
	}

	return largest
}

func calculateArea(arr []int) (int, int, int) {
	minIndex := 0
	maxIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[minIndex] {
			minIndex = i
		}
		if arr[i] >= arr[maxIndex] {
			maxIndex = i
		}
	}
	return arr[minIndex] * len(arr), minIndex, maxIndex
}
