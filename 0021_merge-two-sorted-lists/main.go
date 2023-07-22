package main

import (
	"fmt"
)

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
}

func ReverseList(head *ListNode) (newHead *ListNode) {
	newHead = head
	curr := head.Next
	newHead.Next = nil

	for curr != nil {
		tmp := curr.Next
		curr.Next = newHead
		newHead = curr
		curr = tmp
	}

	return
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (head *ListNode) {
	var curr *ListNode

	for list1 != nil || list2 != nil {
		var tmp *ListNode

		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				tmp = list1
				list1 = list1.Next
			} else {
				tmp = list2
				list2 = list2.Next
			}
		} else if list1 != nil {
			tmp = list1
			list1 = list1.Next
		} else {
			tmp = list2
			list2 = list2.Next
		}

		if head == nil {
			head = tmp
		} else {
			curr.Next = tmp
		}

		curr = tmp
	}

	return
}

func main() {
	l1 := NewList([]int{1, 2, 4})
	l2 := NewList([]int{1, 3, 4, 6})

	l3 := mergeTwoLists(l1, l2)
	PrintList(l3)
}
