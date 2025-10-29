package main

import (
	"github.com/Bychin/leetcode-solutions/utils/linkedlist"
)

func swapNodes(head *linkedlist.ListNode, k int) *linkedlist.ListNode {
	start := &linkedlist.ListNode{
		Next: head,
	}
	slow, fast := start, start

	var (
		left, leftPrev   *linkedlist.ListNode
		right, rightPrev *linkedlist.ListNode
	)
	for range k {
		leftPrev = fast

		fast = fast.Next

		left = fast
	}
	for fast != nil {
		rightPrev = slow

		slow = slow.Next
		fast = fast.Next

		right = slow
	}

	leftPrev.Next, rightPrev.Next = right, left
	left.Next, right.Next = right.Next, left.Next

	return start.Next
}

func main() {
	linkedlist.PrintList(swapNodes(linkedlist.NewList([]int{1, 2, 3, 4, 5}), 2))                // [1,4,3,2,5]
	linkedlist.PrintList(swapNodes(linkedlist.NewList([]int{1, 2, 3, 4, 5}), 3))                // [1,2,3,4,5]
	linkedlist.PrintList(swapNodes(linkedlist.NewList([]int{1, 2, 3, 4, 5, 6}), 3))             // [1,2,4,3,5,6]
	linkedlist.PrintList(swapNodes(linkedlist.NewList([]int{1, 2, 3, 4, 5, 6}), 4))             // [1,2,4,3,5,6]
	linkedlist.PrintList(swapNodes(linkedlist.NewList([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5}), 5)) // [7,9,6,6,8,7,3,0,9,5]
}
