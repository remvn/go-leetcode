package main

import (
	"fmt"
	"math"
	"sort"
)

type TestCase struct {
	arr    []int
	target int
}

func main() {
	tests := []TestCase{
		{
			arr:    []int{-1, 2, 1, -4},
			target: 1,
		},
		{
			arr:    []int{0, 0, 0},
			target: 1,
		},
		{
			arr:    []int{4, 0, 5, -5, 3, 3, 0, -4, -5},
			target: -2,
		},
		{
			arr:    []int{13, 252, -87, -431, -148, 387, -290, 572, -311, -721, 222, 673, 538, 919, 483, -128, -518, 7, -36, -840, 233, -184, -541, 522, -162, 127, -935, -397, 761, 903, -217, 543, 906, -503, -826, -342, 599, -726, 960, -235, 436, -91, -511, -793, -658, -143, -524, -609, -728, -734, 273, -19, -10, 630, -294, -453, 149, -581, -405, 984, 154, -968, 623, -631, 384, -825, 308, 779, -7, 617, 221, 394, 151, -282, 472, 332, -5, -509, 611, -116, 113, 672, -497, -182, 307, -592, 925, 766, -62, 237, -8, 789, 318, -314, -792, -632, -781, 375, 939, -304, -149, 544, -742, 663, 484, 802, 616, 501, -269, -458, -763, -950, -390, -816, 683, -219, 381, 478, -129, 602, -931, 128, 502, 508, -565, -243, -695, -943, -987, -692, 346, -13, -225, -740, -441, -112, 658, 855, -531, 542, 839, 795, -664, 404, -844, -164, -709, 167, 953, -941, -848, 211, -75, 792, -208, 569, -647, -714, -76, -603, -852, -665, -897, -627, 123, -177, -35, -519, -241, -711, -74, 420, -2, -101, 715, 708, 256, -307, 466, -602, -636, 990, 857, 70, 590, -4, 610, -151, 196, -981, 385, -689, -617, 827, 360, -959, -289, 620, 933, -522, 597, -667, -882, 524, 181, -854, 275, -600, 453, -942, 134},
			target: -2805,
		},
	}
	for _, test := range tests {
		fmt.Printf("arr: %v, target: %v \n", test.arr, test.target)
		fmt.Println(threeSumClosest(test.arr, test.target))
	}
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func threeSumClosest(nums []int, target int) int {
	resultDistance := 0
	resultSum := 0
	first := false
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			}
			distance := sum - target
			if abs(distance) < abs(resultDistance) || !first {
				// fmt.Printf("%v+%v+%v = %v \n", nums[left], nums[right], nums[i], sum)
				resultDistance = distance
				resultSum = sum
				first = true
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return resultSum
}
