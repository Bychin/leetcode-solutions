package main

import (
	"github.com/Bychin/leetcode-solutions/utils/linkedlist"
)

func findNextDifferent(head *linkedlist.ListNode) *linkedlist.ListNode {
	val := head.Val
	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val != val {
			return curr
		}
	}
	return nil
}

func deleteDuplicates(head *linkedlist.ListNode) *linkedlist.ListNode {
	start := &linkedlist.ListNode{
		Next: head,
	}

	for prev, curr := start, head; curr != nil; {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			next := findNextDifferent(curr)
			prev.Next = next
			curr = next
		} else {
			prev = curr
			curr = curr.Next
		}
	}

	return start.Next
}

func main() {
	linkedlist.PrintList(deleteDuplicates(linkedlist.NewList([]int{1, 2, 3, 3, 4, 4, 5}))) // [1,2,5]
	linkedlist.PrintList(deleteDuplicates(linkedlist.NewList([]int{1, 1, 1, 2, 3})))       // [2,3]
}
