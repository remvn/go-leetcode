package main

import (
	"fmt"
	"slices"
)

func main() {
	tests := [][]int{
		{1, 2, -3, 3, 1},
		{1, 2, 3, -3, 4},
		{1, 2, 3, -3, -2},
		{1, -1},
		{1, 3, 2, -3, -2, 5, 5, -5, 1},
	}

	for _, test := range tests {
		head := makeList(test)
		fmt.Println(test)
		head = removeZeroSumSublists(head)
		printList(head)
		fmt.Println()
	}
}

func makeList(arr []int) *ListNode {
	curr := &ListNode{
		Val: arr[0],
	}
	head := curr
	for i := 1; i < len(arr); i++ {
		curr.Next = &ListNode{
			Val: arr[i],
		}
		curr = curr.Next
	}
	return head
}

func printList(head *ListNode) {
	fmt.Print("list: ")
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	curr := head
	arr := []*ListNode{}

	index := 0
	for curr != nil {
		arr = append(arr, curr)
		// fmt.Println(arr)
		total := 0
		for i := len(arr) - 1; i >= 0; i-- {
			total += arr[i].Val
			if total == 0 {
				if i-1 < 0 {
					head = curr.Next
				} else {
					// fmt.Println(index, len(arr), i)
					arr[i-1].Next = curr.Next
				}
				arr = slices.Delete(arr, i, len(arr))
				break
			}
		}

		index++
		curr = curr.Next
	}

	return head
}
