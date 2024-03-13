package main

import (
	"fmt"
)

func main() {
	tests := []int{
		8,
		1,
		4,
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(pivotInteger(test))
		fmt.Println()
	}
}

func pivotInteger(n int) int {
	total := n * (n + 1) / 2
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		if sum == total-sum+i {
			return i
		}
	}

	return -1
}

func pivotIntegerTwoPointer(n int) int {
	left := 1
	right := n
	leftSum, rightSum := 0, 0
	for left < right {
		if leftSum <= rightSum {
			leftSum += left
			left++
		} else {
			rightSum += right
			right--
		}
	}

	if leftSum == rightSum {
		return left
	}

	return -1
}
