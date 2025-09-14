package main

import (
	"fmt"
	"strconv"
)

type stack []int

func (s *stack) Push(value int) {
	*s = append(*s, value)
}

func (s *stack) Pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func add(v1, v2 int) int {
	return v1 + v2
}

func sub(v1, v2 int) int {
	return v1 - v2
}

func mul(v1, v2 int) int {
	return v1 * v2
}

func div(v1, v2 int) int {
	return v1 / v2
}

var operators = map[string]func(v1, v2 int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func evalRPN(tokens []string) int {
	s := stack{}
	for _, token := range tokens {
		if fn, ok := operators[token]; ok {
			// note order
			v2 := s.Pop()
			v1 := s.Pop()
			s.Push(fn(v1, v2))
			continue
		}
		v, _ := strconv.Atoi(token)
		s.Push(v)
	}

	return s.Pop()
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))                                             // 9
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))                                            // 6
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})) // 22
}
