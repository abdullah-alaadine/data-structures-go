package linkedlist

import "fmt"

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head   *node
	length int
}

func NewLinkedList() *linkedlist {
	return &linkedlist{}
}

func (l *linkedlist) Prepend(n int) {
	newNode := &node{data: n, next: l.head}
	l.head = newNode
	l.length++
}

func (l *linkedlist) Delete(n int) {
	if l.head == nil {
		return
	}
	if l.head.data == n {
		l.head = l.head.next
		return
	}

	curr := l.head
	for curr.next != nil {
		if curr.next.data == n {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}

func (l *linkedlist) PrintList() {
	curr := l.head
	fmt.Print("[ ")
	for curr != nil {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
	}
	fmt.Println("]")
}
