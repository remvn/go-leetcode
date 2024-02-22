package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, num := range nums {
		look := target - num
		lookIndex, ok := numMap[look]
		if !ok {
			numMap[num] = index
			continue
		}
		return []int{lookIndex, index}
	}
	return []int{}
}
