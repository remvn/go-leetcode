package main

import "fmt"

func main() {
	arr := []string{
		"(abcd)",
		"(u(love)i)",
		"(ed(et(oc))el)",
	}
	for _, str := range arr {
		fmt.Println(reverseParentheses(str))
	}
}

type Pair struct {
	start, end int
}

func reverseParentheses(s string) string {
	pairArr := []Pair{}
	pairStack := []Pair{}
	runes := []rune(s)

	for i := 0; i < len(s); i++ {
		if runes[i] == '(' {
			pairStack = append(pairStack, Pair{start: i})
		}
		if runes[i] == ')' {
			pairStack[len(pairStack)-1].end = i
			pairArr = append(pairArr, pairStack[len(pairStack)-1])
			pairStack = pairStack[:len(pairStack)-1]
		}
	}

	for _, pair := range pairArr {
		left := pair.start
		right := pair.end
		for left < right {
			if runes[left] == '(' {
				left++
				continue
			}
			if runes[right] == ')' {
				right--
				continue
			}
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}

	result := []rune{}
	for _, char := range runes {
		if char == '(' || char == ')' {
			continue
		}
		result = append(result, char)
	}

	return string(result)
}
