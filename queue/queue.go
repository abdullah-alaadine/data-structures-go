package queue

import linkedlist "github.com/abdullah-alaadine/data-structures-go/linked_list"

type Queue struct {
	list linkedlist.Linkedlist
}

func NewQueue() *Queue {
	return &Queue{
		list: *linkedlist.NewLinkedList(),
	}
}

func (q *Queue) Insert(n int) {
	q.list.Append(n)
}

func (q *Queue) Pop() *int {
	return q.list.Shift()
}
