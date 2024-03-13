package main

import "fmt"

func main() {
	tests := []int{0, 1, 2, 3, 4}
	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(climbStairs(test))
		fmt.Println()
	}
}

// see explain: https://leetcode.com/problems/climbing-stairs/description/comments/1570512
func climbStairs(n int) int {
	arr := make([]int, n+3)
	arr[1] = 1
	arr[2] = 2
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}
