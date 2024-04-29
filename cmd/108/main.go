package main

import "fmt"

func main() {
	arr := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(arr)
	printTree(root)
	fmt.Println()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[len(nums)/2]}
	root.Left = sortedArrayToBST(nums[:len(nums)/2])
	root.Right = sortedArrayToBST(nums[len(nums)/2+1:])

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
