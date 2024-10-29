package linkedlist

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func NewNode[T any](v T) *Node[T] {
	return &Node[T]{
		Value: v,
	}
}

type DNode[T any] struct {
	Value T
	Prev  *DNode[T]
	Next  *DNode[T]
}

func NewDNode[T any](v T) *DNode[T] {
	return &DNode[T]{
		Value: v,
	}
}
