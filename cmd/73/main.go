package main

func main() {

}

func setZeroes(matrix [][]int) {
	columnMap := map[int]bool{}
	rowMap := map[int]bool{}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rowMap[i] = true
				columnMap[j] = true
			}
		}
	}

	for col := range columnMap {
		for i := 0; i < len(matrix); i++ {
			matrix[i][col] = 0
		}
	}
	for row := range rowMap {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[row][j] = 0
		}
	}
}
