package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func symmetricNodes(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if p.Left != nil && q.Right == nil || p.Left == nil && q.Right != nil {
		return false
	}

	if p.Right != nil && q.Left == nil || p.Right == nil && q.Left != nil {
		return false
	}

	return true
}

func isSymmetricHelper(p *TreeNode, q *TreeNode) bool {
	if !symmetricNodes(p, q) {
		return false
	}

	if p == nil {
		return true
	}

	if p.Left == nil {
		return isSymmetricHelper(p.Right, q.Left)
	}

	if p.Right == nil {
		return isSymmetricHelper(p.Left, q.Right)
	}

	return isSymmetricHelper(p.Right, q.Left) && isSymmetricHelper(p.Left, q.Right)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	if root.Left == nil || root.Right == nil {
		return false
	}

	return isSymmetricHelper(root.Left, root.Right)
}

func main() {
}
