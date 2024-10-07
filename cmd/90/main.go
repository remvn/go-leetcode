package main

import "slices"

func main() {

}

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	slices.Sort(nums)
	backtrack([]int{}, nums, &result)
	return result
}

func backtrack(arr []int, nums []int, result *[][]int) {
	*result = append(*result, slices.Clone(arr))
	for i, num := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		arr = append(arr, num)
		backtrack(arr, nums[i+1:], result)
		arr = arr[:len(arr)-1]
	}
}
