package main

import (
	"fmt"
)

type Stack []rune

func (s *Stack) Push(c rune) {
	*s = append(*s, c)
}

func (s *Stack) Pop() (rune, error) {
	if len(*s) == 0 {
		return 'a', fmt.Errorf("empty")
	}

	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return value, nil
}

var opened = map[rune]struct{}{
	'(': {},
	'[': {},
	'{': {},
}

var closed = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	stack := Stack{}
	for _, c := range s {
		if _, ok := opened[c]; ok {
			stack.Push(c)
			continue
		}
		if openedC, ok := closed[c]; ok {
			v, err := stack.Pop()
			if err != nil {
				return false
			}
			if openedC != v {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid(""))
}
