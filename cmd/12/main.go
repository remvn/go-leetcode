package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	tests := []int{
		3,
		2,
		9,
		58,
		1994,
	}
	for _, test := range tests {
		fmt.Printf("case %d:\n", test)
		fmt.Println(intToRoman(test))
	}
}

var roman = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func intToRoman(num int) string {
	arr := []string{}
	length := 0
	for num > 0 {
		digit := num % 10
		num = num / 10
		arr = append([]string{digitToRoman(digit, length)}, arr...)
		length++
	}

	strBuilder := strings.Builder{}
	for _, str := range arr {
		strBuilder.WriteString(str)
	}

	return strBuilder.String()
}

func digitToRoman(num int, length int) string {
	pow := int(math.Pow10(length))
	result := roman[num*pow]
	if result != "" {
		return result
	}
	if num > 5 {
		result += roman[5*pow]
		result += strings.Repeat(roman[pow], num-5)
	} else {
		result += strings.Repeat(roman[pow], num)
	}
	return result
}

// Other people's solution. I may use this in production...
func otherIntToRoman(num int) string {
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""
	for num > 0 {
		for i := range value {
			if num >= value[i] {
				result += symbol[i]
				num -= value[i]
				break
			}
		}
	}
	return result
}
