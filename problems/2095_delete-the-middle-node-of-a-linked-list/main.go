package main

import (
	"github.com/Bychin/leetcode-solutions/utils/linkedlist"
)

func deleteMiddle(head *linkedlist.ListNode) *linkedlist.ListNode {
	start := &linkedlist.ListNode{
		Next: head,
	}

	var (
		fast     = head
		slow     = head
		prevSlow = start
	)

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prevSlow = slow
		slow = slow.Next
	}

	prevSlow.Next = slow.Next

	return start.Next
}

func main() {
	linkedlist.PrintList(deleteMiddle(linkedlist.NewList([]int{1, 3, 4, 7, 1, 2, 6}))) // [1,3,4,1,2,6]
	linkedlist.PrintList(deleteMiddle(linkedlist.NewList([]int{1, 2, 3, 4})))          // [1,2,4]
	linkedlist.PrintList(deleteMiddle(linkedlist.NewList([]int{2, 1})))                // [2]
	linkedlist.PrintList(deleteMiddle(linkedlist.NewList([]int{2})))                   // []
}
