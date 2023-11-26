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

func (l *linkedlist) PrintList() {
	curr := l.head
	fmt.Print("[ ")
	for curr != nil {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
	}
	fmt.Println("]")
}
