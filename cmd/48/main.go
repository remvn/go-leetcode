package main

import "fmt"

func main() {
	tests := [][][]int{
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
	}

	for _, matrix := range tests {
		printMatrix(matrix)
		rotate(matrix)
		printMatrix(matrix)
		fmt.Println()
	}
}

func rotate(matrix [][]int) {
	mLen := len(matrix)

	// tranpose the matrix
	for i := 0; i < mLen; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// printMatrix(matrix)

	// reverse each row
	for i := 0; i < mLen; i++ {
		for j := 0; j < mLen/2; j++ {
			matrix[i][j], matrix[i][mLen-j-1] = matrix[i][mLen-j-1], matrix[i][j]
		}
	}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println()
}
