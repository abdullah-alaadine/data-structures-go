package doublylinkedlist

type DoublyLinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length uint
}

func New[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (l *DoublyLinkedList[T]) Len() uint { return l.length }