package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
}

func reverse(x int) int {
	result := 0
	for x != 0 {
		pop := x % 10
		result = result*10 + pop
		x /= 10

		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}
	}

	return result
}
