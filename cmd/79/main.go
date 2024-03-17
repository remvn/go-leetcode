package main

import "fmt"

func main() {
	tests := []TestCase{
		{
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCCED",
		},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(exist(test.board, test.word))
		fmt.Println()
	}
}

type TestCase struct {
	board [][]byte
	word  string
}

func exist(board [][]byte, word string) bool {
	rowSize := len(board)
	colSize := len(board[0])

	var backtrack func(int, int, int) bool
	backtrack = func(row, col, i int) bool {
		if row < 0 || row >= rowSize || col < 0 || col >= colSize || board[row][col] == '0' {
			return false
		}

		if board[row][col] != word[i] {
			return false
		}

		if i == len(word)-1 {
			return true
		}

		originalVal := board[row][col]
		board[row][col] = '0'
		printBoard(board)
		if backtrack(row+1, col, i+1) ||
			backtrack(row, col+1, i+1) ||
			backtrack(row-1, col, i+1) ||
			backtrack(row, col-1, i+1) {
			return true
		}
		board[row][col] = originalVal

		return false
	}

	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			if backtrack(row, col, 0) {
				return true
			}
		}
	}

	return false
}

func printBoard(board [][]byte) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			fmt.Printf("%s ", string(board[row][col]))
		}
		fmt.Println()
	}
	fmt.Println()
}
