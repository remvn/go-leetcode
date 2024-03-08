package main

import "fmt"

func main() {
	tests := [][]int{
		{1, 2, 2, 3, 1, 4},
		{1, 2, 3, 4, 5},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(maxFrequencyElements(test))
		fmt.Println()
	}
}

func maxFrequencyElements(nums []int) int {
	frequency := map[int]int{}
	maxFrequency := 0
	for _, val := range nums {
		frequency[val] += 1
		if frequency[val] > maxFrequency {
			maxFrequency = frequency[val]
		}
	}

	total := 0
	for _, val := range frequency {
		if val == maxFrequency {
			total += val
		}
	}
	return total
}
