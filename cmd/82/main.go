package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 3, 3, 4, 4, 5, 5, 5}
	root := makeListFromArr(arr)
	printList(root)
	root = deleteDuplicates(root)
	printList(root)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 2 2 2 3 4 5
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	countMap := make(map[int]int)
	temp := head
	for temp != nil {
		countMap[temp.Val] += 1
		temp = temp.Next
	}

	if countMap[head.Val] > 1 {
		for head != nil {
			if countMap[head.Val] > 1 {
				head = head.Next
			} else {
				break
			}
		}
	}

	if head == nil {
		return nil
	}

	slow := head
	fast := head.Next
	for fast != nil {
		// fmt.Println(slow.Val, fast.Val)
		if countMap[fast.Val] > 1 {
			slow.Next = fast.Next
		} else {
			slow = slow.Next
		}
		fast = fast.Next
	}

	return head
}

func printList(root *ListNode) {
	for root != nil {
		fmt.Printf("%d ", root.Val)
		root = root.Next
	}
	fmt.Println()
}

func makeListFromArr(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	root := &ListNode{
		Val: arr[0],
	}

	temp := root
	for i := 1; i < len(arr); i++ {
		temp.Next = &ListNode{
			Val: arr[i],
		}
		temp = temp.Next
	}

	return root
}
