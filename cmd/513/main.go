package main

import "fmt"

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	printTree(root.Left)
	fmt.Printf("%d ", root.Val)
	printTree(root.Right)
}

func findBottomLeftValue(root *TreeNode) int {
	arr := []int{}
	getLevelArr(root, height(root), &arr)
	return arr[0]
}

func getLevelArr(root *TreeNode, level int, arr *[]int) {
	if root == nil {
		return
	}

	if level > 1 {
		getLevelArr(root.Left, level-1, arr)
		getLevelArr(root.Right, level-1, arr)
		return
	}

	if level == 1 {
		*arr = append(*arr, root.Val)
	}
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := height(root.Left)
	right := height(root.Right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
