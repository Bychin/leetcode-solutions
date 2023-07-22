package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := helper(root.Left)
	r := helper(root.Right)

	if l == -1 || r == -1 {
		return -1
	}

	diff := abs(l - r)
	if diff > 1 {
		return -1
	}

	return max(l, r) + 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBalanced(root *TreeNode) bool {
	return helper(root) != -1
}

func main() {
}
