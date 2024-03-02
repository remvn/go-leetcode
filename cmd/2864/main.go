package main

import (
	"fmt"
	"strings"
)

func main() {
	tests := []string{
		"010",
		"0101",
	}
	for _, test := range tests {
		fmt.Printf("test: %s result: %s \n", test, maximumOddBinaryNumber(test))
	}
}

func maximumOddBinaryNumber(s string) string {
	count1 := countChar(s, '1')
	if count1 == 1 {
		return strings.Repeat("0", len(s)-1) + "1"
	} else {
		return strings.Repeat("1", count1-1) + strings.Repeat("0", len(s)-count1) + "1"
	}
}

func countChar(s string, char rune) int {
	count := 0
	for _, r := range s {
		if r == char {
			count += 1
		}
	}
	return count
}
