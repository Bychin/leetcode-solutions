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

	fmt.Println()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	var prev *ListNode
	extra := false

	for l1 != nil || l2 != nil || extra {
		if prev != nil {
			curr = &ListNode{}
			prev.Next = curr
		}

		num1, num2 := 0, 0

		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}

		currSum := num1 + num2
		if extra {
			currSum++
			extra = false
		}

		curr.Val = currSum % 10
		if currSum >= 10 {
			extra = true
		}

		prev = curr
	}

	return head
}

func main() {
	l1 := NewList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := NewList([]int{9, 9, 9, 9})
	PrintList(l1)
	PrintList(l2)

	PrintList(addTwoNumbers(l1, l2))
}
