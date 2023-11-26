package linkedlist

import (
	"fmt"
	"os"
)

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

func (l *linkedlist) Insert(n, idx int) {
	if idx > l.length || idx < 0 {
		fmt.Fprintln(os.Stderr, "index out of range")
		return
	}

	if l.head == nil {
		l.head = &node{data: n}
		l.length++
		return
	}

	curr := l.head
	for ; idx > 1; idx-- {
		curr = curr.next
	}
	newNode := &node{data: n, next: curr.next}
	curr.next = newNode
	l.length++
}

func (l *linkedlist) Append(n int) {
	newNode := &node{data: n}
	if l.head == nil {
		l.head = newNode
		l.length++
		return
	}
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
	l.length++
}

func (l *linkedlist) PutIn(n, idx int) {
	if idx >= l.length || idx < 0 {
		fmt.Fprintln(os.Stderr, "index out of range")
		return
	}

	curr := l.head
	for ; idx > 0; idx-- {
		curr = curr.next
	}
	curr.data = n
}

func (l *linkedlist) Delete(n int) {
	if l.head == nil {
		return
	}
	if l.head.data == n {
		l.head = l.head.next
		l.length--
		return
	}

	curr := l.head
	for curr.next != nil {
		if curr.next.data == n {
			curr.next = curr.next.next
			l.length--
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
