package main

import "fmt"

func main() {
	tests := []int{
		4,
		8,
	}

	for _, test := range tests {
		fmt.Printf("sqrt of %d: %d\n", test, mySqrt(test))
		fmt.Println()
	}
}

func mySqrt(x int) int {
	left := 0
	right := x
	// 0 1 2 3 4 5 6 7 8
	for left <= right {
		mid := (left + right) / 2
		// fmt.Println(left, right, mid)
		midSquare := mid * mid
		if midSquare == x {
			return mid
		}
		if midSquare > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}
