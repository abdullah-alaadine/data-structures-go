package stack

import (
	linkedlist "github.com/abdullah-alaadine/data-structures-go/linked_list"
)

type Stack struct {
	list linkedlist.Linkedlist
}

func NewStack() *Stack {
	return &Stack{
		list: *linkedlist.NewLinkedList(),
	}
}

func (s *Stack) Push(v int) {
	s.list.Prepend(v)
}

func (s *Stack) Pop() *int {
	return s.list.Shift()
}

func (s *Stack) Length() int {
	return s.list.Length()
}

func (s *Stack) PrintStack() {
	s.list.PrintList()
}
