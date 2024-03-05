package main

import "fmt"

func main() {
	tests := []string{
		"ca",
		"cac",
		"cabaabac",
		"aabccabba",
		"bbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbccbcbcbccbbabbb",
		"bbbbbbbbbbbbbbbbbbb",
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(minimumLength(test))
		fmt.Println()
	}
}

func minimumLength(s string) int {
	// ca
	// cabaabac
	// aabccabba
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			break
		}
		// fmt.Println(s[left : right+1])
		left++
		right--
		for left < len(s) && s[left] == s[left-1] {
			left++
		}
		for right >= 0 && s[right] == s[right+1] {
			right--
		}
	}

	// fmt.Printf("left: %d, right: %d\n", left, right)

	if left > right {
		return 0
	}
	// fmt.Printf("result: %s\n", s[left:right+1])

	return len(s[left : right+1])
}
