package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back == nil {
		return nil
	}
	return s.v.Remove(back)
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func main() {
	stack := NewStack()
	for i := 1; i < 5; i++ {
		stack.Push(i)
	}

	val := stack.Pop()
	for val != nil {
		fmt.Printf("%v -> ", val)
		val = stack.Pop()
	}
}
