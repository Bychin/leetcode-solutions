package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) (result int) {
	for head != nil {
		result = result<<1 | head.Val
		head = head.Next
	}
	return result
}

func main() {
	// fmt.Println(getDecimalValue([1,0,1]])) // 5
}
