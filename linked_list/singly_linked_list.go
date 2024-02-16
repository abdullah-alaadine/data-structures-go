package linkedlist

import (
	"fmt"
)

type LinkedList[T any] struct {
	head   *Node[T]
	length int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
	}
}

func (l *LinkedList[T]) Len() int { return l.length }

func (l *LinkedList[T]) Prepend(v T) {
	newNode := NewNode[T](v)
	l.prepend(newNode)
}

func (l *LinkedList[T]) prepend(n *Node[T]) {
	n.Next = l.head
	l.head = n
	l.length++
}

func (l *LinkedList[T]) Append(v T) {
	newNode := NewNode[T](v)
	l.append(newNode)
}

func (l *LinkedList[T]) append(n *Node[T]) {
	if l.length == 0 {
		l.head = n
		l.length++
		return
	}
	curr := l.head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = n
	l.length++
}

func (l *LinkedList[T]) Remove(index int) (T, bool) {
	return l.remove(index)
}

func (l *LinkedList[T]) Pop() (T, bool) {
	return l.remove(0)
}

func (l *LinkedList[T]) RemoveTail() (T, bool) {
	return l.remove(l.length - 1)
}

func (l *LinkedList[T]) remove(index int) (T, bool) {
	var t T
	if index < 0 || index >= l.length {
		return t, false
	}
	if index == 0 {
		temp := l.head.Value
		l.head = l.head.Next
		l.length--
		return temp, true
	}
	curr := l.head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}
	temp := curr.Next.Value
	curr.Next = curr.Next.Next
	l.length--
	return temp, true
}

func (l *LinkedList[T]) PrintList() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%v ", curr.Value)
		curr = curr.Next
	}
	fmt.Print("\n")
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.length == 0
}

func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.length = 0
}
