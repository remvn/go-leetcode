package main

import "fmt"

func main() {
	tests := []TestCase{
		{3, 7},
		{3, 2},
	}

	for _, test := range tests {
		fmt.Printf("m: %d, n: %d\n", test.m, test.n)
		fmt.Println(uniquePaths(test.m, test.n))
		fmt.Println()
	}
}

type TestCase struct {
	m, n int
}

// solved using Pascal triangle
// see: https://leetcode.com/problems/unique-paths/description/comments/2041560
func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := 0
			if i == 0 || j == 0 {
				sum = 1
			} else {
				sum = matrix[i-1][j] + matrix[i][j-1]
			}
			matrix[i][j] = sum
		}
	}

	return matrix[m-1][n-1]
}
