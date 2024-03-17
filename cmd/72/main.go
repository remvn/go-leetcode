package main

import (
	"fmt"
)

func main() {
	tests := []TestCase{
		{"horse", "ros"},
		// {"intention", "execution"},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(minDistance(test.word1, test.word2))
		fmt.Println()
	}
}

type TestCase struct {
	word1, word2 string
}

// see https://www.youtube.com/watch?v=XYi2-LPrwm4
func minDistance(word1 string, word2 string) int {
	cached := make([][]int, len(word1)+1)
	for i := 0; i < len(cached); i++ {
		cached[i] = make([]int, len(word2)+1)
	}

	for i := 0; i < len(word1)+1; i++ {
		cached[i][len(word2)] = len(word1) - i
	}
	for i := 0; i < len(word2)+1; i++ {
		cached[len(word1)][i] = len(word2) - i
	}

	for i := len(word1) - 1; i >= 0; i-- {
		for j := len(word2) - 1; j >= 0; j-- {
			if word1[i] != word2[j] {
				cached[i][j] = 1 + min(cached[i+1][j+1], cached[i+1][j], cached[i][j+1])
			} else {
				cached[i][j] = cached[i+1][j+1]
			}
		}
	}

	return cached[0][0]
}
