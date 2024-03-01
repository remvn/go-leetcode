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

	var prev, curr, newHead *ListNode = nil, head, head.Next
	index := 0
	for curr != nil && curr.Next != nil {
		if index%2 == 0 {
			firstNode, secondNode := curr, curr.Next

			firstNode.Next = secondNode.Next
			secondNode.Next = firstNode

			if prev != nil {
				prev.Next = secondNode
			}
			curr = secondNode
		}
		index++
		prev = curr
		curr = curr.Next
	}
	return newHead
}
