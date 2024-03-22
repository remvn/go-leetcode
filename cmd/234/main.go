package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}

	stack := make([]int, 0, length)
	curr = head
	index := 0
	for curr != nil {
		if !(length%2 == 1 && index == length/2) {
			if index >= length/2 && curr.Val == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, curr.Val)
			}
		}

		curr = curr.Next
		index++
	}

	return len(stack) == 0
}
