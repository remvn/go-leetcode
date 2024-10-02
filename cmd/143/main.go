package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	head := makeList(arr)
	printLinkedList(head)
	reorderList(head)
	printLinkedList(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(arr []int) *ListNode {
	head := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	temp := head
	for i := 1; i < len(arr); i++ {
		temp.Next = &ListNode{
			Val: arr[i],
		}
		temp = temp.Next
	}

	return head
}

// 1 2 3 4 5
// 1 5 2 4 3

// 1 2 3 4
// 1 4 2 3

func reorderList(head *ListNode) {
	stack := []*ListNode{}
	temp := head.Next
	for temp != nil {
		stack = append(stack, temp)
		temp = temp.Next
	}

	temp = head
	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if temp == pop { // list length is odd
			temp.Next = nil
			break
		}
		if temp.Next == pop { // list length is even
			pop.Next = nil
			break
		}
		prev := temp.Next // 2
		temp.Next = pop   // 1 -> 5
		pop.Next = prev   // 5 -> 2
		temp = prev       // temp = 2
	}
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
