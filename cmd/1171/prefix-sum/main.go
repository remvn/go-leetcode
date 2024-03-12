package main

import "fmt"

func main() {
	tests := [][]int{
		{1, 2, -3, 3, 1},
		{1, 2, 3, -3, 4},
		{1, 2, 3, -3, -2},
		{1, -1},
		{1, 3, 2, -3, -2, 5, 5, -5, 1},
		{-1, 1, 0, 1},
	}

	for _, test := range tests {
		head := makeList(test)
		printList(head)
		head = removeZeroSumSublists(head)
		printList(head)
		fmt.Println()
	}
}

func makeList(arr []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for i := 0; i < len(arr); i++ {
		curr.Next = &ListNode{
			Val: arr[i],
		}
		curr = curr.Next
	}
	return dummy.Next
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
	prefixMap := map[int]*ListNode{}

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	prefixSum := 0
	for curr := dummy; curr != nil; curr = curr.Next {
		prefixSum += curr.Val

		if prev, found := prefixMap[prefixSum]; found {
			toDelete := prev.Next
			sum := prefixSum
			if toDelete != nil {
				sum += toDelete.Val
			}

			// Go back to the point where the sum sequence
			// started. Remove items from map until we reach
			// the prefixSum again
			for toDelete != nil && sum != prefixSum {
				delete(prefixMap, sum)
				toDelete = toDelete.Next
				if toDelete != nil {
					sum += toDelete.Val
				}
			}

			prev.Next = curr.Next
		} else {
			prefixMap[prefixSum] = curr
		}
	}

	return dummy.Next
}
