package main

import (
	"fmt"
	"slices"
)

func main() {
	tests := []TestCase{
		{
			order: "cba",
			s:     "abcd",
		},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(customSortString(test.order, test.s))
		fmt.Println()
	}
}

type TestCase struct {
	order string
	s     string
}

func customSortString(order string, s string) string {
	orderMap := map[rune]int{}
	orderRunes := []rune(order)
	for i := 0; i < len(orderRunes); i++ {
		orderMap[orderRunes[i]] = i
	}

	match := []int{}
	remain := make([]rune, 0, len(s))
	for _, char := range s {
		if index, ok := orderMap[char]; ok {
			match = append(match, index)
		} else {
			remain = append(remain, char)
		}
	}

	slices.Sort(match)
	result := make([]rune, 0, len(s))
	for _, i := range match {
		result = append(result, orderRunes[i])
	}
	result = append(result, remain...)

	return string(result)
}
