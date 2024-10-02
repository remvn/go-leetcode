package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	length := removeDuplicates(arr)
	fmt.Println(arr[:length])
}

// 1 1 1 2 2 3
// 1 1 2 2 3

// i + tripLength;

func removeDuplicates(nums []int) int {
	length := len(nums)
	count := 1
	for i := 1; i < length; i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		if count == 3 {
			count = 1
			trimLength := 1
			for j := i + 1; j < length; j++ {
				if nums[j] == nums[i] {
					trimLength++
				} else {
					break
				}
			}

			for j := i; j < length; j++ {
				if j+trimLength >= length {
					break
				}
				nums[j] = nums[j+trimLength]
			}

			length -= trimLength
		}
	}

	return length
}
