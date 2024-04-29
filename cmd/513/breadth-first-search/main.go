package main

import (
	"fmt"

	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
)

func main() {
	arr := []int{
		1, 2, 3, 4, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 5, 6, -1, -1, 7,
	}
	root := createTree(arr)
	printTree(root)
	fmt.Println()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTree(arr []int) *TreeNode {
	root := &TreeNode{Val: arr[0]}
	i := 1
	queue := []*TreeNode{root}
	next_left := true
	for i < len(arr) && len(queue) > 0 {
		node := queue[0]
		if arr[i] >= 0 {
			if next_left {
				node.Left = &TreeNode{Val: arr[i]}
				queue = append(queue, node.Left)
			} else {
				node.Right = &TreeNode{Val: arr[i]}
				queue = append(queue, node.Right)
			}
		}
		if !next_left {
			queue = queue[1:]
		}
		next_left = !next_left
		i++
	}
	return root
}

func printTree(node *TreeNode) {
	if node == nil {
		return
	}

	printTree(node.Left)
	fmt.Printf("%d ", node.Val)
	printTree(node.Right)
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
