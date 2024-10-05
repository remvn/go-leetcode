package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	arr := []int{5, 6, 9, 1, 2, 6}
	printHeap(arr)
	heapify(arr)
	printHeap(arr)

	sorted := []int{}
	for len(arr) > 0 {
		num, newArr := pop(arr)
		arr = newArr
		sorted = append(sorted, num)
	}
	fmt.Println(sorted)
}

func heapify(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		shiftDown(arr, i)
	}
}

func pop(arr []int) (int, []int) {
	if len(arr) == 0 {
		return 0, arr
	}
	num := arr[0]
	arr[0] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	shiftDown(arr, 0)
	return num, arr
}

func shiftDown(arr []int, index int) {
	maxIndex := index
	n := len(arr)
	left := index*2 + 1
	right := index*2 + 2

	if left < n && arr[left] > arr[maxIndex] {
		maxIndex = left
	}

	if right < n && arr[right] > arr[maxIndex] {
		maxIndex = right
	}

	if maxIndex != index {
		arr[maxIndex], arr[index] = arr[index], arr[maxIndex]
		shiftDown(arr, maxIndex)
	}
}

func printSpaces(count int) {
	if count > 0 {
		fmt.Print(strings.Repeat(" ", count))
	}
}

// 1 2 3 4 5
//   1
// 2   3
//4 5 6 7

func printHeap(arr []int) {
	arrLen := len(arr)
	height := int(math.Floor(math.Log2(float64(arrLen)))) + 1 // Calculate the height of the tree
	canvasSize := int(math.Pow(2, float64(height))) - 1

	levelStart := 0 // Start of the level
	prevSize := 1
	for i := 0; i < height; i++ {
		levelSize := int(math.Pow(2, float64(i)))
		levelEnd := levelStart + levelSize

		printSpaces(canvasSize / levelSize / 2)
		for j := levelStart; j < levelEnd && j < arrLen; j++ {
			fmt.Printf("%d", arr[j])
			if i > 0 {
				printSpaces(canvasSize / prevSize / 2)
			}
		}
		printSpaces(canvasSize / levelSize / 2)

		fmt.Println()

		levelStart = levelEnd
		prevSize = levelSize
	}
}

func findRelativeRanks(score []int) []string {
	medalMap := make(map[int]string)

	temp := slices.Clone(score)
	heapify(temp)
	count := 0
	for len(temp) > 0 {
		count++
		num, newTemp := pop(temp)
		temp = newTemp

		switch count {
		case 1:
			medalMap[num] = "Gold Medal"
		case 2:
			medalMap[num] = "Silver Medal"
		case 3:
			medalMap[num] = "Bronze Medal"
		default:
			medalMap[num] = strconv.Itoa(count)
		}
	}

	medalArr := []string{}
	for _, num := range score {
		medalArr = append(medalArr, medalMap[num])
	}

	return medalArr
}
