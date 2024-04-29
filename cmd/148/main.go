package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 5, 6, 0, 7, 3, 4}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}

	pivotIndex := left - 1
	pivotVal := nums[right]
	for i := left; i <= right; i++ {
		if nums[i] <= pivotVal {
			pivotIndex++
			nums[i], nums[pivotIndex] = nums[pivotIndex], nums[i]
		}
	}

	quickSort(nums, pivotIndex+1, right)
	quickSort(nums, left, pivotIndex-1)
}
