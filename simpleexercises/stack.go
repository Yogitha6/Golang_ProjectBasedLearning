package main

import (
	"fmt"
)

type node struct {
	value int
	bottom *node
}

type stack struct {
	top *node
}

func (s *stack) push(value int) {
	if s.top==nil {
		n := node {value: value, bottom: nil}
		s.top = &n
	} else {
		n := node {value: value, bottom: s.top}
		s.top = &n
	}
}

func (s *stack) pop() int {
	if s.top == nil {
		fmt.Println("No element to Pop")
		return -1
	} else {
		temp := s.top.value
		s.top = s.top.bottom
		return temp
	}
}

func (s *stack) peek() int {
	if s.top == nil {
		fmt.Println("No element to Peek")
		return -1
	} else {
		return s.top.value
	}
}

func (s *stack) isEmpty() bool {
	if s.top == nil {
		return true
	}
	return false
}

func (s *stack) size() int {
	size := 0
	temp:= s.top
	for temp!=nil {
		temp = temp.bottom
		size = size + 1
	}
	return size
}

func main() {
	fmt.Println("Stack Implementation")
	s := stack{}
	s.push(1)
	s.push(4)
	s.push(5)
	fmt.Println(s.pop())
	s.push(6)
	fmt.Println(s.pop())
	fmt.Println(s.size())
	fmt.Println(s.isEmpty())
	fmt.Println(s.peek())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.isEmpty())
}
