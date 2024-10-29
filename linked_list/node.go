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
