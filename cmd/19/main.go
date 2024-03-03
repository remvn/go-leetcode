package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := lengthOfList(head)
	if length == 1 {
		return nil
	}
	if length-n-1 < 0 {
		return head.Next
	}

	curr := head
	index := 0
	for curr != nil {
		if index == length-n-1 {
			break
		}
		index++
		curr = curr.Next
	}
	if curr != nil && curr.Next != nil {
		curr.Next = curr.Next.Next
	}
	return head
}

func lengthOfList(head *ListNode) int {
	length := 0
	for head != nil {
		head = head.Next
		length++
	}
	return length
}
