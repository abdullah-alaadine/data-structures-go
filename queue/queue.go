package queue

import linkedlist "github.com/knbr13/data-structures-go/linked_list"

type Queue[T any] struct {
	list linkedlist.LinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		list: *linkedlist.New[T](),
	}
}

func (q *Queue[T]) Len() int { return q.list.Len() }

func (q *Queue[T]) Push(v T) {
	q.list.Append(v)
}

func (q *Queue[T]) Pop() (T, bool) {
	return q.list.Remove(q.list.Len() - 1)
}

func (q *Queue[T]) Clear() {
	q.list.Clear()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}
