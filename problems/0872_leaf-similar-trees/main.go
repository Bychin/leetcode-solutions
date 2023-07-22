package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}

	for i := 1; i < len(nums); i++ {
		root.Append(nums[i])
	}

	return root
}

func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d\n", root.Val)
	fmt.Printf("---(%d) left call---\n", root.Val)
	PrintTree(root.Left)

	fmt.Printf("---(%d) right call---\n", root.Val)
	PrintTree(root.Right)
}

func (root *TreeNode) Append(num int) {
	if root == nil {
		return
	}

	if root.Val <= num {
		if root.Right != nil {
			root.Right.Append(num)
			return
		}

		root.Right = &TreeNode{
			Val: num,
		}
		return
	}

	// left branch

	if root.Left != nil {
		root.Left.Append(num)
		return
	}

	root.Left = &TreeNode{
		Val: num,
	}
}

func (root *TreeNode) Insert(num int) {
	if root == nil {
		return
	}

	if root.Val <= num {
		if root.Right == nil {
			root.Right = &TreeNode{
				Val: num,
			}
			return
		}

		if root.Right.Val <= num {
			root.Right.Insert(num)
			return
		}

		tmp := root.Right
		root.Right = &TreeNode{
			Val:   num,
			Right: tmp,
		}
		return
	}

	// left branch

	if root.Left == nil {
		root.Left = &TreeNode{
			Val: num,
		}
		return
	}

	if root.Left.Val > num {
		root.Left.Insert(num)
		return
	}

	tmp := root.Left
	root.Left = &TreeNode{
		Val:  num,
		Left: tmp,
	}
}

// DFS

func rangeSumBST(root *TreeNode, low int, high int) (sum int) {
	if root == nil {
		return 0
	}

	if root.Val >= low && root.Val <= high {
		sum += root.Val
		sum += rangeSumBST(root.Left, low, high)
		sum += rangeSumBST(root.Right, low, high)
		return
	}

	if root.Val > high {
		sum += rangeSumBST(root.Left, low, high)
		return
	}

	sum += rangeSumBST(root.Right, low, high)
	return
}

func main() {
	tree1 := NewBST([]int{10, 5, 15, 3, 7, 18})
	PrintTree(tree1)

	fmt.Println(rangeSumBST(tree1, 7, 15))
}
