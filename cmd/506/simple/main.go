package main

import (
	"slices"
	"strconv"
)

func main() {
}

func findRelativeRanks(score []int) []string {
	temp := []int{}
	medalMap := make(map[int]string)

	temp = append(temp, score...)
	slices.SortFunc(temp, func(a, b int) int {
		switch {
		case a == b:
			return 0
		case a > b:
			return -1
		default:
			return 1
		}
	})

	for index, num := range temp {
		medal := ""
		switch index {
		case 0:
			medal = "Gold Medal"
		case 1:
			medal = "Silver Medal"
		case 2:
			medal = "Bronze Medal"
		default:
			medal = strconv.Itoa(index + 1)
		}
		medalMap[num] = medal
	}

	medalArr := []string{}
	for _, num := range score {
		medalArr = append(medalArr, medalMap[num])
	}

	return medalArr
}
