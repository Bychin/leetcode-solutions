package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListWithCycle(nums []int, pos int) (head *ListNode) {
	if len(nums) == 0 {
		return nil
	}

	var prev *ListNode
	var cycleTo *ListNode

	for i, n := range nums {
		new := &ListNode{
			Val: n,
		}

		if prev == nil {
			prev = new
			head = new

			if i == pos {
				cycleTo = new
			}

			continue
		}

		prev.Next = new
		prev = new

		if i == pos {
			cycleTo = new
		}
	}

	if cycleTo != nil {
		prev.Next = cycleTo
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
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for {
		if slow == nil || fast == nil {
			return false
		}
		if slow == fast {
			return true
		}

		slow = slow.Next
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
	}
}

func main() {
	l1 := NewListWithCycle([]int{1, 2, 3, 4}, 3)
	PrintList(l1)

	fmt.Println(hasCycle(l1))
}
