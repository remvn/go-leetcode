package main

import "fmt"

func main() {
	tests := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
	}
	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(maxArea(test))
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	maxArea := 0
	left := 0
	right := len(height) - 1
	for left < right {
		cHeight := min(height[left], height[right])
		area := cHeight * (right - left)
		if area > maxArea {
			maxArea = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}
