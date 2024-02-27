package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	tests := []testCase{
		{arr: []int{1, 0, -1, 0, -2, 2}, target: 0},
		{arr: []int{2, 2, 2, 2, 2}, target: 8},
	}

	for _, test := range tests {
		fmt.Printf("arr: %+v target: %v \n", test.arr, test.target)
		fmt.Printf("result: %+v \n", fourSum(test.arr, test.target))
	}
}

type testCase struct {
	arr    []int
	target int
}

// After a search through other solutions
// I found that we can use a map[[4]int]bool (dint even know that we can do this...)
// see: https://stackoverflow.com/questions/20297503/slice-as-a-key-in-map
// to get rid of skipping a, b (very tricky part)
// then this problem is very simple

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	// fmt.Println(nums)

	result := [][]int{}

	for a := 0; a < len(nums)-3; a++ { // 0 1 2 3 4
		for b := a + 1; b < len(nums)-2; b++ {
			c := b + 1
			d := len(nums) - 1
			for c < d {
				total := nums[a] + nums[b] + nums[c] + nums[d]
				// fmt.Printf("%d %d %d %d \n", a, b, c, d)
				// fmt.Printf("%d %d %d %d = %d \n", nums[a], nums[b], nums[c], nums[d], total)
				if total == target {
					result = append(result, []int{nums[a], nums[b], nums[c], nums[d]})
					c++
					d--

					// tricky part:
					// We skip duplicate numbers until the end of duplicated sequence.
					// * This only affect those 2 outer loop and
					// the rest of the "for c < d" loop stay the same
					//   2 2 2 3 3 3 (before skipping)
					//   | |
					//   a b
					//
					// All the match for 2 2 x x has been found
					//   2 2 2 3 3 3 (after skipping),
					//       |
					//      a,b
					//
					//   2 2 2 3 3 3 (2 outer loop run, a++, b=a+1)
					//         | |
					//         a b
					for a+1 < len(nums) && nums[a+1] == nums[a] {
						a++
					}
					for b+1 < len(nums) && nums[b+1] == nums[b] {
						b++
					}

					// The part below is like problem 15 (3 sum)
					for c < d && nums[c] == nums[c-1] {
						c++
					}
					for d > c && nums[d] == nums[d+1] {
						d--
					}
				} else {
					if total < target {
						c++
					} else {
						d--
					}
				}
			}
		}
	}

	return result
}
