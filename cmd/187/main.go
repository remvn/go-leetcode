package main

import "fmt"

func main() {
	arr := []string{
		"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
		"AAAAAAAAAAAAA",
		"AAAAAAAAAAA",
	}

	for _, s := range arr {
		fmt.Println(findRepeatedDnaSequences(s))
	}
}

func findRepeatedDnaSequences(s string) []string {
	countMap := make(map[string]int)
	for i := 0; i+9 < len(s); i++ {
		seq := s[i : i+10]
		countMap[seq] += 1
	}

	arr := []string{}
	for seq, count := range countMap {
		if count > 1 {
			arr = append(arr, seq)
		}
	}

	return arr
}
