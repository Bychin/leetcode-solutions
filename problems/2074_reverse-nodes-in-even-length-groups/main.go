package main

import (
	"github.com/Bychin/leetcode-solutions/utils/linkedlist"
)

func reverseList(start, end *linkedlist.ListNode) *linkedlist.ListNode {
	var prev *linkedlist.ListNode

	for curr := start; curr != end; {
		tmp := curr.Next
		curr.Next = prev

		prev, curr = curr, tmp
	}

	return prev
}

func reverseEvenLengthGroups(head *linkedlist.ListNode) *linkedlist.ListNode {
	curr, prevEnd := head, head
	for groupLen := 1; groupLen <= 10e5 && curr != nil; groupLen++ {
		groupStart, groupEnd := curr, curr
		var nextStart *linkedlist.ListNode

		// find the end of a group and the actual len
		actualLen := 0
		for j := 0; j < groupLen && curr != nil; j++ {
			actualLen++
			groupEnd = curr
			curr = curr.Next
			nextStart = curr
		}

		// skip odd group
		if actualLen%2 != 0 {
			prevEnd = groupEnd
			continue
		}

		reversedGroup := reverseList(groupStart, nextStart)
		prevEnd.Next = reversedGroup
		if nextStart != nil {
			groupStart.Next = nextStart // after reverse groupStart is the last element in the group
			prevEnd = groupStart
			curr = nextStart
		}
	}

	return head
}

func main() {
	linkedlist.PrintList(reverseEvenLengthGroups(linkedlist.NewList([]int{5, 2, 6, 3, 9, 1, 7, 3, 8, 4}))) // [5,6,2,3,9,1,4,8,3,7]
	linkedlist.PrintList(reverseEvenLengthGroups(linkedlist.NewList([]int{1, 1, 0, 6})))                   // [1,0,1,6]
	linkedlist.PrintList(reverseEvenLengthGroups(linkedlist.NewList([]int{1, 1, 0, 6, 5})))                // [1,0,1,5,6]
}
