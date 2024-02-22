package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(input, numRows))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([][]rune, numRows)
	direction := -1
	rowIndex := 0

	for _, char := range s {
		if rowIndex+1 >= numRows || rowIndex == 0 {
			direction *= -1
		}
		rows[rowIndex] = append(rows[rowIndex], char)
		rowIndex += direction
	}

	// merge rows
	result := strings.Builder{}
	for _, row := range rows {
		for _, r := range row {
			result.WriteRune(r)
		}
	}

	return result.String()
}
