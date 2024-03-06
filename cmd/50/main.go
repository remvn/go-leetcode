package main

import "fmt"

func main() {
	tests := []TestCase{
		{x: 2, n: 10},
		{x: 2.10000, n: 3},
		{x: 2.00000, n: -2},
		{x: -1, n: 2147483647},
	}

	for _, test := range tests {
		fmt.Printf("x: %f n: %d\n", test.x, test.n)
		fmt.Printf("x^n = %f\n", myPow(test.x, test.n))
		fmt.Println()
	}
}

type TestCase struct {
	x float64
	n int
}

// TODO: nah.. I will check how to do bit manipulation when I have time
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n *= -1
		x = 1 / x
	}

	result := float64(1)
	for n > 0 {
		result *= x
		n--
	}

	return result
}
