package doublylinkedlist

type DoublyLinkedList[T any] struct {
	head   *DNode[T]
	tail   *DNode[T]
	length uint
}
