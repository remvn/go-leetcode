package main

func main() {

}

var roman = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && roman[s[i]] < roman[s[i+1]] {
			total -= roman[s[i]]
		} else {
			total += roman[s[i]]
		}
	}
	return total
}
