package queue

import linkedlist "github.com/knbr13/data-structures-go/linked_list"

type Queue[T any] struct {
	list linkedlist.LinkedList[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		list: *linkedlist.New[T](),
	}
}

func (q *Queue[T]) Len() int { return q.list.Len() }

func (q *Queue[T]) Push(v T) {
	q.list.Append(v)
}

func (q *Queue[T]) Pop() (T, bool) {
	return q.list.Pop()
}

func (q *Queue[T]) Clear() {
	q.list.Clear()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue[T]) PrintQueue() {
	q.list.PrintList()
}
