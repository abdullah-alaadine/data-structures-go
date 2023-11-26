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

func (l *linkedlist) prepend(n int) {
	newNode := &node{data: n, next: l.head}
	l.head = newNode
	l.length++
}

func (l *linkedlist) printList() {
	curr := l.head
	fmt.Print("[ ")
	for curr != nil {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
	}
	fmt.Println("]")
}
