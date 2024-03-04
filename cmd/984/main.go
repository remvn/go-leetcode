package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := []TestCase{
		{arr: []int{100, 200, 300, 400}, power: 200},
		{arr: []int{100}, power: 50},
		{arr: []int{200, 100}, power: 150},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(bagOfTokensScore(test.arr, test.power))
		fmt.Println()
	}
}

type TestCase struct {
	arr   []int
	power int
}

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	// keep spending power until we run out of them
	// then we will increase by playing token at the
	// end of sorted array (largest)

	score := 0
	left := 0
	right := len(tokens) - 1
	for left <= right {
		// fmt.Printf("1. power: %d, score: %d \n", power, score)
		if power >= tokens[left] {
			power -= tokens[left]
			left++
			score++
		} else if score > 0 && left != right {
			// left == right means this is the last move, dont play to save score
			power += tokens[right]
			right--
			score--
		} else {
			// out of score, out of power
			break
		}
		// fmt.Printf("2. power: %d, score: %d \n", power, score)
		// fmt.Println()
	}

	return score
}
