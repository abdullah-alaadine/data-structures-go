package queue

import linkedlist "github.com/knbr13/data-structures-go/linked_list"

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

func (q *Queue) Len() int { return q.list.Length() }

func (q *Queue) PrintQueue() {
	q.list.PrintList()
}
