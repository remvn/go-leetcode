package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	head := makeList(arr)
	printList(head)
	head = reverseList(head)
	printList(head)
}

func makeList(nums []int) *ListNode {
	head := &ListNode{}
	curr := head
	for _, num := range nums {
		curr.Next = &ListNode{
			Val: num,
		}
		curr = curr.Next
	}
	return head.Next
}

func printList(head *ListNode) {
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

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		next := head.Next
		if prev != nil {
			head.Next = prev
		} else {
			head.Next = nil
		}
		prev = head
		head = next
	}

	return prev
}
