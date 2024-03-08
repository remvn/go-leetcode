package main

import (
	"fmt"
	"strings"
)

func main() {
	tests := []string{
		"Hello World",
		"   fly me   to   the moon  ",
		"luffy is still joyboy",
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(lengthOfLastWord(test))
		fmt.Println()
	}
}

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	return len(arr[len(arr)-1])
}
