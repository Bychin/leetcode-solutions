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

func CompareUntilNil(one *ListNode, another *ListNode) bool {
	for one != nil && another != nil {
		if one.Val != another.Val {
			return false
		}

		one, another = one.Next, another.Next
	}

	return true
}

func isPalindrome(head *ListNode) bool {
	slowPointer := head
	fastPointer := head

	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}

	// now slow pointer is in the middle

	return CompareUntilNil(head, ReverseList(slowPointer))
}

func main() {
	fmt.Println(isPalindrome(NewList([]int{1, 2, 2, 1})))
}
