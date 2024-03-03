package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	nextPermutation(arr)
	fmt.Println(arr)
}

func nextPermutation(nums []int) {
	// if the whole array is already largest
	if isLargest(nums) {
		sort(nums)
		return
	}

	for i := len(nums) - 1; i >= 0; i-- {
		subArr := nums[i:]
		// fmt.Println(subArr)
		// sliding i until subArr isn't already largest
		if !isLargest(subArr) {

			firstDigit := subArr[0]

			// find the smallest number but larger than first digit
			// it's still larger but as small as possbile
			smallestIndex := 1
			for j := 1; j < len(subArr); j++ {
				if subArr[j] < subArr[smallestIndex] && subArr[j] > firstDigit {
					smallestIndex = j
				}
			}
			// swap
			subArr[0], subArr[smallestIndex] = subArr[smallestIndex], subArr[0]

			// minimize further by sorting remaining digits
			sort(subArr[1:])

			break
		}
	}
}

func sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[index] {
				index = j
			}
		}
		// swap
		arr[i], arr[index] = arr[index], arr[i]
	}
}

func isLargest(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] < nums[i] {
			return false
		}
	}
	return true
}

/*
	Explanation from leetcode
	We want to find a larger number, but not too large - just the next larger one =>
	We want to avoid chaning the number in the left - it will increase the number too much =>
	Example: 4325413 -> we can only change the last two numbers and have 432531
	What if it was 432531 in the first place? 31 cannot be increased.
	Lets try 531 - still no
	2531 - this can be incrased - the smallest number that can be used to increase the 2 is 3. so for now we have 3521.
	Next we want to minimize 3521 - thats easier - just sort the numbers to the right of 3 - 3125. So the unswer is 4323125
*/
