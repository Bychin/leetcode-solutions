package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (l *MyLinkedList) Get(index int) int {
	curr := l.head
	for i := 0; i < index; i++ {
		if curr == nil {
			return -1
		}
		curr = curr.Next
	}
	if curr != nil {
		return curr.Val
	} else {
		return -1
	}
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.head = &Node{
		Val:  val,
		Next: l.head,
	}
}

func (l *MyLinkedList) AddAtTail(val int) {
	if l.head == nil {
		l.AddAtHead(val)
		return
	}
	curr := l.head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = &Node{
		Val: val,
	}
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		l.AddAtHead(val)
		return
	}
	prev := &Node{
		Next: l.head,
	}
	curr := l.head
	for i := 0; i < index; i++ {
		if curr == nil {
			return
		}
		prev = curr
		curr = curr.Next
	}
	prev.Next = &Node{
		Val:  val,
		Next: prev.Next,
	}
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if l.head != nil {
			l.head = l.head.Next
		}
		return
	}
	prev := &Node{
		Next: l.head,
	}
	curr := l.head
	for i := 0; i < index; i++ {
		if curr == nil {
			return
		}
		prev = curr
		curr = curr.Next
	}
	if curr != nil {
		prev.Next = curr.Next
	} else {
		prev.Next = nil
	}
}

func PrintList(l *MyLinkedList) {
	curr := l.head

	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

func main() {
	l := Constructor()
	l.AddAtHead(2)
	l.DeleteAtIndex(1)
	l.AddAtHead(2)
	l.AddAtHead(3)
	PrintList(&l)
}
