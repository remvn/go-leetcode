package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 2, 5, 0, 1, 8, 7, 6, 9, 4}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func bucketSort(nums []int) {
	bucket := map[int]int{}
	for _, num := range nums {
		bucket[num]++
	}

	offset := 0
	for _, num := range []int{0, 1, 2} {
		for count := 0; count < bucket[num]; count++ {
			nums[offset] = num
			offset++
		}
	}
}

func quickSort(nums []int, left int, right int) {
	if left < right {
		pivot := partition(nums, left, right)
		quickSort(nums, pivot+1, right)
		quickSort(nums, left, pivot-1)
	}
}

func partition(nums []int, left int, right int) int {
	pivot := nums[right]
	index := left - 1
	for i := left; i <= right; i++ {
		if nums[i] <= pivot {
			index++
			swap(&nums[index], &nums[i])
		}
	}
	// swap(&nums[index], &nums[right])
	return index
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
