package queue

import (
	linkedlist "github.com/abdullah-alaadine/data-structures-go/linked_list"
)

type Queue linkedlist.Linkedlist

func NewQueue() *Queue {
	return (*Queue)(linkedlist.NewLinkedList())
}
