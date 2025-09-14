package main

import "fmt"

type stackNode struct {
	value int
	min   int
}

type MinStack []stackNode

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) topNode() stackNode {
	return (*m)[len(*m)-1]
}

func (m *MinStack) Push(value int) {
	min := value
	if len(*m) != 0 {
		topNodeMin := m.topNode().min
		if topNodeMin < value {
			min = topNodeMin
		}
	}

	*m = append(*m, stackNode{
		value: value,
		min:   min,
	})
}

func (m *MinStack) Pop() {
	(*m) = (*m)[:len(*m)-1]
}

func (m *MinStack) Top() int {
	return m.topNode().value
}

func (m *MinStack) GetMin() int {
	return m.topNode().min
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2
}
