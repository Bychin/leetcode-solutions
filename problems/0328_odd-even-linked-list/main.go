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

func appendToList(list *ListNode, node *ListNode) *ListNode {
	tmp1 := node.Next
	tmp2 := list.Next

	list.Next = node
	node.Next = tmp2

	return tmp1
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var oddRunner, evenRunner *ListNode
	var oddStart, evenStart *ListNode

	oddFlag := true
	for head != nil {
		if oddFlag {
			oddFlag = false

			if oddRunner == nil {
				oddRunner = head
				oddStart = oddRunner

				head = head.Next
				continue
			}

			head = appendToList(oddRunner, head)
			oddRunner = oddRunner.Next
			continue
		}

		oddFlag = true

		if evenRunner == nil {
			evenRunner = head
			evenStart = head

			head = head.Next
			continue
		}

		head = appendToList(evenRunner, head)
		evenRunner = evenRunner.Next
	}

	oddRunner.Next = evenStart
	evenRunner.Next = nil

	return oddStart
}

func main() {
	PrintList(oddEvenList(NewList([]int{1, 2, 3, 4, 5})))
	PrintList(oddEvenList(NewList([]int{2, 1, 3, 5, 6, 4, 7})))
	PrintList(oddEvenList(NewList([]int{2})))
	PrintList(oddEvenList(NewList(nil)))
}
