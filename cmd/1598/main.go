package main

import "fmt"

func main() {
	testCases := [][]string{
		{"d1/", "d2/", "../", "d21/", "./"},
		{"d1/", "d2/", "./", "d3/", "../", "d31/"},
		{"d1/", "../", "../", "../"},
	}
	for _, logs := range testCases {
		fmt.Println(minOperations(logs))
	}
}

func minOperations(logs []string) int {
	level := 0
	for _, log := range logs {
		switch log {
		case "../":
			if level > 0 {
				level--
			}
		case "./":
		default:
			level++
		}
	}

	return level
}
