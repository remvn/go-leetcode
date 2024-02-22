package main

import "fmt"

func main() {
	arr1 := []int{2, 4, 3}
	arr2 := []int{5, 6, 4}
	l1 := createList(arr1)
	l2 := createList(arr2)
	result := addTwoNumbers(l1, l2)
	printList(result)
}

func createList(arr []int) *ListNode {
	list := new(ListNode)
	curr := list
	for i := 0; i < len(arr); i++ {
		curr.Val = arr[i]
		if i+1 < len(arr) {
			curr.Next = new(ListNode)
		}
		curr = curr.Next
	}
	return list
}

func printList(node *ListNode) {
	curr := node
	for curr != nil {
		fmt.Printf("%v ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	sum, carry := 0, 0
	current := result

	for l1 != nil || l2 != nil || carry != 0 {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		val := sum % 10
		carry = sum / 10
		current.Next = &ListNode{
			Val: val,
		}
		current = current.Next
	}
	return result.Next
}
