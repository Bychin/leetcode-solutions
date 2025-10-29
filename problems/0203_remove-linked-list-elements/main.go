package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var prev *ListNode

	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val != val {
			prev = curr
			continue
		}
		if prev == nil { // we are at head
			head = curr.Next
		} else {
			prev.Next = curr.Next
		}
	}
	return head
}

func main() {
	// fmt.Println(removeElements([1,2,6,3,4,5,6], 6)) // [1,2,3,4,5]
	// fmt.Println(removeElements([1,2,2,1], 2)) // [1,1]
}
