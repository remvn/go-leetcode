package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	length := getLength(head)
	loop := k % length

	tail, pre := getTail(head)
	for i := 0; i < loop; i++ {
		tail.Next = head
		head = tail
		pre.Next = nil
		tail, pre = getTail(head)
	}

	return head
}

func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

func getTail(head *ListNode) (*ListNode, *ListNode) {
	tail := head
	pre := head
	for tail.Next != nil {
		pre = tail
		tail = tail.Next
	}
	return tail, pre
}
