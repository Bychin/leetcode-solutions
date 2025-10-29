package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := head
	fast := head
	// create offset using fast pointer
	for i := 1; fast != nil && i <= n; i++ {
		fast = fast.Next
	}
	// move slow pointer to the element to be deleted
	var prevSlow *ListNode
	for fast != nil {
		fast = fast.Next
		prevSlow = slow
		slow = slow.Next
	}
	// delete element
	if slow != nil {
		if prevSlow == nil {
			head = slow.Next
		} else {
			prevSlow.Next = slow.Next
		}
	}
	return head
}

func main() {
	// fmt.Println(removeNthFromEnd([1,2,3,4,5], 2)) // [1,2,3,5]
}
