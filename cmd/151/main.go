package main

import (
	"fmt"
	"strings"
)

func main() {
	tests := []string{
		"the sky is blue",
		"  hello world  ",
		"a good   example",
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(reverseWords(test))
		fmt.Println()
	}
}

func reverseWords(s string) string {
	arr := strings.Fields(s)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return strings.Join(arr, " ")
}
