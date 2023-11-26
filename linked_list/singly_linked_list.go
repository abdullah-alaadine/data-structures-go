package linkedlist

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
}
