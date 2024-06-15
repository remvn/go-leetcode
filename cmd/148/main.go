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

	// -1 at the start and increment before swapping
	// to avoid 1 leftover
	pivotIndex := left - 1
	pivotVal := nums[right]
	for i := left; i <= right; i++ {
		if nums[i] <= pivotVal {
			pivotIndex++
			nums[i], nums[pivotIndex] = nums[pivotIndex], nums[i]
		}
	}
	// must swap the last element in order to make left part smaller than nums[right]
	// and right part is larger than nums[right]
	//
	// we must do this below if for and if dont have `=`
	// pivotIndex++
	// nums[pivotIndex], nums[right] = nums[left], nums[pivotIndex]

	quickSort(nums, pivotIndex+1, right)
	quickSort(nums, left, pivotIndex-1)
}
