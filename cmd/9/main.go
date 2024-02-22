package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	origin := x
	rev := 0
	for x != 0 {
		pop := x % 10
		rev = rev*10 + pop
		x = x / 10
	}
	return rev == origin
}
