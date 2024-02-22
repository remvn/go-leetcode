package main

import "fmt"

func main() {
	tests := []string{
		"dvdf",
		"abcabcbb",
		"bbbbb",
		"pwwkew",
	}
	for _, test := range tests {
		fmt.Printf("%s: %v \n", test, lengthOfLongestSubstring(test))
	}
}

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	result := 0

	// - index of arr where it's the longest (i - start + 1 = current length)
	// - update when charMap contains current char (repeating char)
	// - update with charMap[char] + 1
	start := 0

	for i := 0; i < len(s); i++ {
		char := s[i]
		if lastIndex, ok := charMap[char]; ok {
			if lastIndex >= start {
				start = lastIndex + 1
			}
		}
		charMap[char] = i
		if length := i - start + 1; length > result {
			result = length
		}
	}

	return result
}
