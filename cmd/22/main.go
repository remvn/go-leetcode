package main

import "fmt"

func main() {
	arr := generateParenthesis(3)
	fmt.Println(arr)
	fmt.Println(len(arr))
}

func generateParenthesis(n int) []string {
	arr := []string{}
	generate(n*2, "(", &arr)
	validArr := []string{}
	for _, str := range arr {
		if checkValid(str) {
			validArr = append(validArr, str)
		}
	}
	return validArr
}

func generate(length int, str string, arr *[]string) {
	if len(str)+1 == length {
		str += ")"
		*arr = append(*arr, str)
		return
	}
	generate(length, str+"(", arr)
	generate(length, str+")", arr)
}

func checkValid(str string) bool {
	stack := []rune{}
	for _, r := range str {
		if r == '(' {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
