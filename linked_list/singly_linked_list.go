package linkedlist

import (
	"fmt"
	"os"
)

type node struct {
	data int
	next *node
}

type Linkedlist struct {
	head   *node
	length int
}

func NewLinkedList() *Linkedlist {
	return &Linkedlist{}
}

func (l *Linkedlist) Prepend(n int) {
	newNode := &node{data: n, next: l.head}
	l.head = newNode
	l.length++
}

func (l *Linkedlist) Insert(n, idx int) {
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

func (l *Linkedlist) Append(n int) {
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

func (l *Linkedlist) PutIn(n, idx int) {
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

func (l *Linkedlist) Delete(n int) {
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

func (l *Linkedlist) Shift() *int {
	if l.head == nil {
		return nil
	}
	v := l.head
	l.head = l.head.next
	l.length--
	return &v.data
}

func (l *Linkedlist) DeleteAt(idx int) {
	if idx >= l.length || idx < 0 {
		fmt.Fprintln(os.Stderr, "index out of range")
		return
	}

	if idx == 0 {
		l.head = l.head.next
		l.length--
		return
	}

	curr := l.head
	for ; idx > 1; idx-- {
		curr = curr.next
	}
	curr.next = curr.next.next
	l.length--
}

func (l *Linkedlist) PrintList() {
	curr := l.head
	fmt.Print("[ ")
	for curr != nil {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
	}
	fmt.Println("]")
}

func (l *Linkedlist) PrintListReversely() {
	if l.head == nil {
		return
	}

	values := []int{}
	curr := l.head

	for curr != nil {
		values = append(values, curr.data)
		curr = curr.next
	}
	fmt.Print("[ ")
	for i := len(values) - 1; i >= 0; i-- {
		fmt.Printf("%d ", values[i])
	}
	fmt.Println("]")
}

func (l *Linkedlist) Length() int {
	return l.length
}

func (l *Linkedlist) GetHeadData() *int {
	if l.head != nil {
		return &l.head.data
	}
	return nil
}
