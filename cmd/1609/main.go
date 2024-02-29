package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	treeHeight := height(root)
	arr := make([][]int, treeHeight)

	levelArr(root, 0, &arr)
	for levelIndex, levelArr := range arr {
		isEvenIndex := levelIndex%2 == 0
		for i := 0; i < len(levelArr); i++ {
			if isEvenIndex {
				if i > 0 && levelArr[i-1] >= levelArr[i] {
					return false
				}
				if levelArr[i]%2 == 0 {
					return false
				}
			} else {
				if i > 0 && levelArr[i-1] <= levelArr[i] {
					return false
				}
				if levelArr[i]%2 == 1 {
					return false
				}
			}
		}
	}

	return true
}

func levelArr(root *TreeNode, index int, levels *[][]int) {
	if root == nil {
		return
	}

	levelArr(root.Left, index+1, levels)
	levelArr(root.Right, index+1, levels)

	curr := (*levels)[index]
	curr = append(curr, root.Val)
	(*levels)[index] = curr
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := height(root.Left)
	right := height(root.Right)

	return max(left, right) + 1
}
