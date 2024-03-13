package main

import "fmt"

func main() {
	tests := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
		{9, 9},
	}
	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(plusOne(test))
		fmt.Println()
	}
}

func plusOne(digits []int) []int {
	digits[len(digits)-1] += 1
	arr := make([]int, len(digits)+1)
	index := len(arr) - 1
	for i := len(digits) - 1; i >= 0; i-- {
		arr[index] += digits[i]
		temp := arr[index]
		arr[index] = temp % 10
		arr[index-1] += temp / 10
		index--
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr = arr[i:]
			break
		}
	}

	return arr
}
