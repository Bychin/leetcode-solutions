package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func equalNodes(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if p.Left != nil && q.Left == nil || p.Left == nil && q.Left != nil {
		return false
	}

	if p.Right != nil && q.Right == nil || p.Right == nil && q.Right != nil {
		return false
	}

	return true
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if !equalNodes(p, q) {
		return false
	}

	if p == nil {
		return true
	}

	if p.Left == nil {
		return isSameTree(p.Right, q.Right)
	}

	if p.Right == nil {
		return isSameTree(p.Left, q.Left)
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
}
