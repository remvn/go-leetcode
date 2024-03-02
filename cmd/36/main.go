package main

import (
	"strconv"
)

func main() {

}

// find duplicated number in board
// by using matrix of col, row and grid
// to keep track of numbers that already exist
func isValidSudoku(board [][]byte) bool {
	var cols [9][9]bool
	var rows [9][9]bool
	var grids [9][9]bool

	for col := 0; col < 9; col++ {
		for row := 0; row < 9; row++ {
			str := board[col][row]
			val, err := strconv.Atoi(string(str))
			if err != nil {
				continue
			}

			val--
			// gridIndex grow by 1 every 3 col and by 3 every 3 row
			gridIndex := col/3 + (row/3)*3
			if cols[col][val] || rows[row][val] || grids[gridIndex][val] {
				return false
			}

			cols[col][val] = true
			rows[row][val] = true
			grids[gridIndex][val] = true
		}
	}

	return true
}
