package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(multiply("45", "25"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	arr := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			val := int(num1[i]-'0') * int(num2[j]-'0') // digit char in byte - 48 = digit in int
			arr[i+j+1] += val

			arr[i+j] += arr[i+j+1] / 10
			arr[i+j+1] = arr[i+j+1] % 10
		}
	}

	// skip empty space in result array eg: 25*45 = 0,0,1,1,2,5
	index := 0
	for _, num := range arr {
		if num != 0 {
			break
		}
		index++
	}

	result := ""
	for i := index; i < len(arr); i++ {
		digitStr := strconv.Itoa(arr[i])
		result += digitStr
	}

	return result
}
