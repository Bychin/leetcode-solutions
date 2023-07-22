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

	for i := 0; curr != nil; i++ {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
		if i == 8 {
			break
		}
	}

	fmt.Println()
}

func RevertList(root *ListNode) (newList *ListNode) {
	curr := root

	for curr != nil {
		tmp := curr.Next

		curr.Next = newList
		newList = curr

		curr = tmp
	}

	return newList
}

func reverseList(head *ListNode) (newHead *ListNode) {
	if head == nil {
		return nil
	}

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

func main() {
	PrintList(reverseList(NewList([]int{1, 2, 2, 4})))
	PrintList(RevertList(NewList([]int{1, 2, 2, 4})))
	PrintList(RevertList(NewList([]int{1, 2, 3, 4})))
	PrintList(RevertList(NewList([]int{1})))
	PrintList(RevertList(NewList(nil)))
}
