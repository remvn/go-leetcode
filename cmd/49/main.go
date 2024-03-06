package main

import (
	"fmt"
	"slices"
)

func main() {
	tests := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
	}
	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(groupAnagrams(test))
		fmt.Println()
	}
}

func groupAnagrams(strs []string) [][]string {
	mapped := map[string][]string{}
	for _, str := range strs {
		runes := []rune(str)
		slices.Sort(runes)
		sorted := string(runes)
		mapped[sorted] = append(mapped[sorted], str)
	}

	result := [][]string{}
	for _, val := range mapped {
		result = append(result, val)
	}

	return result
}
