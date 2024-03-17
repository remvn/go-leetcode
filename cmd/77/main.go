package main

import "fmt"

func main() {
	tests := []TestCase{
		{4, 2}, {1, 1},
	}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(combine(test.n, test.k))
		fmt.Println()
	}
}

type TestCase struct {
	n, k int
}

func combine(n int, k int) [][]int {
	result := [][]int{}
	backtracking([]int{}, 1, k, n, &result)
	return result
}

func backtracking(arr []int, index int, k int, n int, result *[][]int) {
	// fmt.Println(arr)
	if len(arr) == k {
		*result = append(*result, arr)
		return
	}

	for i := index; i <= n; i++ {
		clone := cloneAndAppend(arr, i, k)
		backtracking(clone, i+1, k, n, result)
	}
}

func cloneAndAppend(arr []int, val int, arrCap int) []int {
	clone := make([]int, len(arr), arrCap)
	copy(clone, arr)
	clone = append(clone, val)
	return clone
}
