package main

import "fmt"

func main() {
	tests := [][]int{
		{1, 2, 3, 4},
		{1, 2, 3},
	}

	for _, test := range tests {
		root := createList(test)
		printList(root)
		root = swapPairs(root)
		printList(root)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(root *ListNode) {
	for root != nil {
		fmt.Printf("%d ", root.Val)
		root = root.Next
	}
	fmt.Println()
}

func createList(arr []int) *ListNode {
	root := &ListNode{
		Val: arr[0],
	}
	curr := root
	for i := 1; i < len(arr); i++ {
		curr.Next = &ListNode{
			Val: arr[i],
		}
		curr = curr.Next
	}
	return root
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	prev := head
	index := 0
	for curr != nil {
		fmt.Println(curr.Val)
		if index%2 == 0 && curr.Next != nil {
			// swap
			swap := curr.Next
			next := swap.Next

			swap.Next = curr
			curr.Next = next

			curr = swap

			if index == 0 {
				head = swap
			} else {
				prev.Next = swap
			}
		}
		index++
		prev = curr
		curr = curr.Next
	}
	return head
}
