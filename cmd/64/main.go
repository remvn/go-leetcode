package main

import (
	"fmt"
	"math"
)

func main() {
	tests := [][][]int{
		{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
		{{1, 2, 3}, {4, 5, 6}},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(minPathSum(test))
		fmt.Println()
	}
}

func minPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if j == 0 && i == 0 {
				continue
			}
			top, left := math.MaxInt, math.MaxInt
			if j-1 >= 0 {
				left = grid[i][j-1]
			}
			if i-1 >= 0 {
				top = grid[i-1][j]
			}
			grid[i][j] += min(top, left)
			// fmt.Printf("%d ", grid[i][j])
		}
		// fmt.Println()
	}

	return grid[len(grid)-1][len(grid[0])-1]
}
