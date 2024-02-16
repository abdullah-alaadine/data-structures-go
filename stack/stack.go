package stack

import (
	linkedlist "github.com/knbr13/data-structures-go/linked_list"
)

type Stack[T any] struct {
	list linkedlist.LinkedList[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		list: *linkedlist.New[T](),
	}
}

func (s *Stack[T]) Push(v T) {
	s.list.Prepend(v)
}

func (s *Stack[T]) Pop() (T, bool) {
	return s.list.Pop()
}

func (s *Stack[T]) Len() int {
	return s.list.Len()
}

func (s *Stack[T]) PrintStack() {
	s.list.PrintList()
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *Stack[T]) Clear() {
	s.list.Clear()
}
