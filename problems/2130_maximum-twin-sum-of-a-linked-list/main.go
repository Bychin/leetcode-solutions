package main

import (
	"fmt"

	"github.com/Bychin/leetcode-solutions/utils/linkedlist"
)

func reverseList(head *linkedlist.ListNode) *linkedlist.ListNode {
	var prev *linkedlist.ListNode
	for curr := head; curr != nil; {
		next := curr.Next
		curr.Next = prev
		curr, prev = next, curr
	}
	return prev
}

func pairSum(head *linkedlist.ListNode) (sum int) {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	secondHead := reverseList(slow)

	for secondHead != nil {
		currSum := head.Val + secondHead.Val
		sum = max(sum, currSum)
		head, secondHead = head.Next, secondHead.Next
	}
	return sum
}

func main() {
	fmt.Println(pairSum(linkedlist.NewList([]int{5, 4, 2, 1}))) // 6
	fmt.Println(pairSum(linkedlist.NewList([]int{4, 2, 2, 3}))) // 7
}
