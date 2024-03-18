package main

import (
	"fmt"
	"slices"
)

func main() {
	tests := [][][]int{
		{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
		{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
		{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(findMinArrowShots(test))
		fmt.Println()
	}
}

func findMinArrowShots(points [][]int) int {
	slices.SortFunc(points, func(a, b []int) int {
		if a[1] == b[1] {
			return 0
		}
		if a[1] < b[1] {
			return -1
		} else {
			return 1
		}
	})

	arrowCount := 1
	arrow := points[0][1]
	for _, point := range points {
		// fmt.Println(point, arrow)
		if arrow >= point[0] && arrow <= point[1] {
			continue
		} else {
			arrowCount++
			arrow = point[1]
		}
	}

	return arrowCount
}
