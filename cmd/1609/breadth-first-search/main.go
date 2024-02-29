package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	q := []*TreeNode{}
	q = append(q, root)

	level := 0
	for len(q) > 0 {
		// previous level's size
		size := len(q)
		prev := -1
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}

			if level%2 == 0 {
				if prev > 0 && prev >= node.Val {
					return false
				}
				if node.Val%2 == 0 {
					return false
				}
			} else {
				if prev > 0 && prev <= node.Val {
					return false
				}
				if node.Val%2 == 1 {
					return false
				}
			}

			prev = node.Val
		}

		level++
	}

	return true
}
