package main

import "fmt"

func main() {
	tests := []string{
		"()",
		"()[]{}",
		"(]",
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(isValid(test))
	}
}

var pairs = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

func isValid(s string) bool {
	stack := []rune{}
	for _, char := range s {
		if _, ok := pairs[char]; ok {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			popChar := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if compare, ok := pairs[popChar]; !ok || compare != char {
				return false
			}
		}
	}

	return len(stack) == 0
}
