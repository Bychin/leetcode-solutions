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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}

	pivot := len(lists) / 2

	l := mergeKLists(lists[:pivot])
	r := mergeKLists(lists[pivot:])

	return mergeTwoLists(l, r)
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
	l1 := NewList([]int{1, 4, 5})
	l2 := NewList([]int{1, 3, 4})
	l3 := NewList([]int{2, 6})

	l4 := mergeKLists([]*ListNode{l1, l2, l3})
	PrintList(l4)
}
