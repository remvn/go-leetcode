package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := generateTrees(3)
	for _, tree := range arr {
		printTree(tree)
		fmt.Println()
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	treeMap := make(map[int][]*TreeNode)
	for index := range nums {
		num := nums[index]
		root := &TreeNode{
			Val: num,
		}
		newNums := slices.Concat(nums[:index], nums[index+1:])
		generateNode(root, newNums, treeMap)
	}

	nodeArr := []*TreeNode{}
	for _, arr := range treeMap {
		nodeArr = append(nodeArr, arr...)
	}
	return nodeArr
}

// 1
// 1 2
// 1 3
// 2
// 2 1
// 2 3

func generateNode(root *TreeNode, nums []int, treeMap map[int][]*TreeNode) {
	if len(nums) == 0 {
		// printTree(root)
		// fmt.Println()
		nodeArr, ok := treeMap[root.Val]
		if ok {
			for _, node := range nodeArr {
				if isTreeEqual(root, node) {
					return
				}
			}
		}
		treeMap[root.Val] = append(treeMap[root.Val], root)
		return
	}

	for index, num := range nums {
		newRoot := &TreeNode{}
		cloneTree(root, newRoot)
		appendToTree(newRoot, num)

		newNums := slices.Concat(nums[:index], nums[index+1:])
		generateNode(newRoot, newNums, treeMap)
	}
}

func isTreeEqual(treeA *TreeNode, treeB *TreeNode) bool {
	queueA := []*TreeNode{treeA}
	queueB := []*TreeNode{treeB}

	enqueue := func(queue *[]*TreeNode, node *TreeNode) {
		if node.Left != nil {
			*queue = append(*queue, node.Left)
		}
		if node.Right != nil {
			*queue = append(*queue, node.Right)
		}
	}

	for len(queueA) > 0 && len(queueB) > 0 {
		nodeA := queueA[0]
		nodeB := queueB[0]
		queueA = queueA[1:]
		queueB = queueB[1:]
		if nodeA.Val != nodeB.Val {
			return false
		}
		enqueue(&queueA, nodeA)
		enqueue(&queueB, nodeB)
	}

	return len(queueA) == len(queueB)
}

func appendToTree(root *TreeNode, value int) {
	var target *TreeNode
	for root != nil {
		target = root
		if value > root.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	if value > target.Val {
		target.Right = &TreeNode{
			Val: value,
		}
	} else {
		target.Left = &TreeNode{
			Val: value,
		}
	}
}

func cloneTree(root *TreeNode, newRoot *TreeNode) {
	if root == nil {
		return
	}

	newRoot.Val = root.Val
	if root.Left != nil {
		newRoot.Left = &TreeNode{
			Val: root.Left.Val,
		}
		cloneTree(root.Left, newRoot.Left)
	}
	if root.Right != nil {
		newRoot.Right = &TreeNode{
			Val: root.Right.Val,
		}
		cloneTree(root.Right, newRoot.Right)
	}
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}
