package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix[0]), len(matrix)
	top, left, right, bottom := -1, -1, m, n // edges
	x, y := 0, 0
	direction := "right"

	result := make([]int, 0, m*n)
	for len(result) < m*n {
		// fmt.Println(matrix[y][x])
		result = append(result, matrix[y][x])
		switch direction {
		case "right":
			if x+1 == right { // check if we hit the edge
				direction = "down"
				top += 1
				y += 1
			} else {
				x += 1
			}
		case "left":
			if x-1 == left {
				direction = "up"
				bottom -= 1
				y -= 1
			} else {
				x -= 1
			}
		case "down":
			if y+1 == bottom {
				direction = "left"
				right -= 1
				x -= 1
			} else {
				y += 1
			}
		case "up":
			if y-1 == top {
				direction = "right"
				left += 1
				x += 1
			} else {
				y -= 1
			}
		}
	}

	return result
}
