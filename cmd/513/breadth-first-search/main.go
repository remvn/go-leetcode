package main

import (
	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
)

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	return breadthFirstSearch(root)
}

func breadthFirstSearch(root *TreeNode) int {
	queue := llq.New()
	queue.Enqueue(root)
	result := -1
	for queue.Size() > 0 {
		item, ok := queue.Dequeue()
		if ok {
			temp := item.(*TreeNode)
			result = temp.Val
			if temp.Right != nil {
				queue.Enqueue(temp.Right)
			}
			if temp.Left != nil {
				queue.Enqueue(temp.Left)
			}
		}
	}
	return result
}
