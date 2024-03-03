package main

import (
	"fmt"
)

func main() {
	digits := []string{"232", "", "2"}
	for _, digit := range digits {
		fmt.Println(digit)
		fmt.Println(letterCombinations(digit))
		fmt.Println()
	}
}

func letterCombinations(digits string) []string {
	result := []string{}
	if digits == "" {
		return result
	}

	keyboard := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var backtracking func(digits string, subset []rune)
	backtracking = func(digits string, subset []rune) {
		if digits == "" {
			result = append(result, string(subset))
		} else {
			letters := keyboard[rune(digits[0])]
			for _, letter := range letters {
				// we could use the code in the comment to
				// have a better understanding
				// subset = append(subset, letter)
				// backtracking(digits[1:], subset)
				// subset = slices.Delete(subset, len(subset)-1, len(subset))
				backtracking(digits[1:], append(subset, letter))
			}
		}
	}

	backtracking(digits, nil)

	return result
}
