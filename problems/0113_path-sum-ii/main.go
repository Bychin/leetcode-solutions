package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, targetSum int, path []int) [][]int {
	if root == nil {
		return nil
	}

	diff := targetSum - root.Val

	if root.Left == nil && root.Right == nil {
		if diff != 0 {
			return nil
		}

		newPath := make([]int, 0, len(path))
		newPath = append(newPath, path...)
		newPath = append(newPath, root.Val)

		return [][]int{newPath}
	}

	path = append(path, root.Val)

	leftPath := helper(root.Left, diff, path)
	rightPath := helper(root.Right, diff, path)

	return append(leftPath, rightPath...)
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	return helper(root, targetSum, nil)
}

func main() {
}
