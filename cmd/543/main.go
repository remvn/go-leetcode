package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	node := buildTreeFromArray(arr, 0)
	printInOrder(node)
	fmt.Println()
	fmt.Println(drawTree(node, 0))
	fmt.Printf("result: %d\n", diameterOfBinaryTree(node))
}

func printInOrder(node *TreeNode) {
	if node != nil {
		printInOrder(node.Left)
		fmt.Printf("%d ", node.Val)
		printInOrder(node.Right)
	}
}

func drawTree(node *TreeNode, depth int) string {
	if node == nil {
		return ""
	}

	result := drawTree(node.Left, depth+1)
	result += strings.Repeat("   ", depth) + fmt.Sprintf("%d\n", node.Val)
	result += drawTree(node.Right, depth+1)

	return result
}

func buildTreeFromArray(nums []int, index int) *TreeNode {
	if index >= len(nums) || nums[index] == math.MinInt64 {
		return nil
	}

	root := &TreeNode{Val: nums[index]}
	root.Left = buildTreeFromArray(nums, 2*index+1)
	root.Right = buildTreeFromArray(nums, 2*index+2)

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var longest int

func diameterOfBinaryTree(root *TreeNode) int {
	longest = 0
	find(root)
	return longest
}

// the longest path between two node
// is sum of max length between two branch
func find(node *TreeNode) int {
	if node == nil {
		return 0
	}
	// find the max length between two branch
	left := find(node.Left)
	right := find(node.Right)

	distance := left + right
	longest = max(longest, distance)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
