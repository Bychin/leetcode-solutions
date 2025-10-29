package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums []int) (head *ListNode) {
	if len(nums) == 0 {
		return nil
	}

	var prev *ListNode

	for _, n := range nums {
		new := &ListNode{
			Val: n,
		}

		if prev == nil {
			prev = new
			head = new
			continue
		}

		prev.Next = new
		prev = new
	}

	return
}

func PrintList(head *ListNode) {
	curr := head

	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}
