package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	root := buildTreeFromArray(nums, 0)
	matrix := printTree(root)
	for _, line := range matrix {
		fmt.Println(strings.Join(line, ""))
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeInfo struct {
	node *TreeNode
	row  int
	col  int
}

func printTree(root *TreeNode) [][]string {
	matrix := [][]string{}

	height := treeHeight(root)
	canvasSize := int(math.Pow(2, float64(height))) - 1
	for range height {
		matrix = append(matrix, slices.Repeat([]string{""}, canvasSize))
	}

	queue := []NodeInfo{{
		node: root,
		row:  0,
		col:  canvasSize / 2,
	}}
	for len(queue) > 0 {
		info := queue[0]
		node := info.node
		queue = queue[1:]

		matrix[info.row][info.col] = fmt.Sprint(node.Val)

		if node.Left != nil {
			queue = append(queue, NodeInfo{
				node: node.Left,
				row:  info.row + 1,
				col:  info.col - int(math.Pow(2, float64(height-1)-float64(info.row)-1)),
				// height is off by one from provided equation
				// 3 - 2^(2-0-1)
			})
		}
		if node.Right != nil {
			queue = append(queue, NodeInfo{
				node: node.Right,
				row:  info.row + 1,
				col:  info.col + int(math.Pow(2, float64(height-1)-float64(info.row)-1)),
				// 3 + 2^(2-0-1)
			})
		}
	}

	return matrix
}

//   1
// 2  3
//4 5

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftH := treeHeight(root.Left) + 1
	rightH := treeHeight(root.Right) + 1

	return max(leftH, rightH)
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
