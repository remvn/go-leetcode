package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = removeNth(head, n)
	return head
}

func removeNth(head *ListNode, n int) *ListNode {
	length := lengthOfList(head)
	if length == 1 {
		return nil
	}
	if length-n-1 < 0 {
		return head.Next
	}

	index := 0
	curr := head
	for curr != nil {
		if index == length-n-1 || length-n-1 < 0 {
			break
		}
		curr = curr.Next
		index++
	}
	if curr.Next != nil {
		curr.Next = curr.Next.Next
	}
	return head
}

func lengthOfList(head *ListNode) int {
	size := 0
	for head != nil {
		head = head.Next
		size++
	}
	return size
}
