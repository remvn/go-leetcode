package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(4))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	str := countAndSay(n - 1)
	fmt.Println("n:", n)
	fmt.Println("str:", str)

	result := ""
	char := str[0]
	count := 1
	for i := 1; i < len(str); i++ {
		if str[i] != char {
			result += strconv.Itoa(count) + string(char)
			char = str[i]
			count = 0
		}
		count++
	}

	if count > 0 {
		result += strconv.Itoa(count) + string(char)
	}

	fmt.Println("result:", result)
	fmt.Println()

	return result
}
