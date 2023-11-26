package queue

import (
	linkedlist "github.com/abdullah-alaadine/data-structures-go/linked_list"
)

type Queue struct {
	list linkedlist.Linkedlist
}

func NewQueue() *Queue {
	return &Queue{
		list: *linkedlist.NewLinkedList(),
	}
}

func (q *Queue) Enqueue(v int) {
	q.list.Prepend(v)
}

func (q *Queue) Dequeue() *int {
	return q.list.Shift()
}

func (q *Queue) Length() int {
	return q.list.Length()
}

func (q *Queue) PrintQueue() {
	q.list.PrintList()
}
