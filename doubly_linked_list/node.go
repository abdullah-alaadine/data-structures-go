package doublylinkedlist

type Node[T any] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

func NewNode[T any](v T) *Node[T] {
	return &Node[T]{
		Value: v,
	}
}
