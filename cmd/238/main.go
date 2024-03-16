package main

import "fmt"

func main() {
	tests := [][]int{
		{1, 2, 3, 4},
		{-1, 1, 0, -3, 3},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(productExceptSelf(test))
		fmt.Println()
	}
}

func productExceptSelf(nums []int) []int {
	zeroCount := 0
	zeroIndex := -1
	totalProduct := 1
	for i, num := range nums {
		if num == 0 {
			zeroCount++
			zeroIndex = i
			if zeroCount == 2 {
				return make([]int, len(nums))
			}
		} else {
			totalProduct *= num
		}
	}

	result := make([]int, len(nums))
	for i, num := range nums {
		if i == zeroIndex {
			result[i] = totalProduct
		} else {
			if zeroCount == 0 {
				result[i] = totalProduct / num
			}
		}
	}

	return result
}
