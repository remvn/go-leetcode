package main

import "fmt"

func main() {
	tests := [][][]int{
		{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		{{0, 1}, {0, 0}},
		{{0, 0}, {1, 1}, {0, 0}},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(uniquePathsWithObstacles(test))
		fmt.Println()
	}
}

// solved using Pascal triangle
// see: https://leetcode.com/problems/unique-paths/description/comments/2041560
// with a minor adjustment:
// if obstacleGrid[i][j] == 1 we "delete" this path by
// set it to zero
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			sum := 0
			if i == 0 && j == 0 {
				sum = 1
			} else {
				if i-1 >= 0 {
					sum += obstacleGrid[i-1][j]
				}
				if j-1 >= 0 {
					sum += obstacleGrid[i][j-1]
				}
			}
			obstacleGrid[i][j] = sum
		}
	}

	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
