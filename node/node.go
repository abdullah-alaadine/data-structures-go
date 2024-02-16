package node

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func New[T any](v T) *Node[T] {
	return &Node[T]{
		Value: v,
		Next:  nil,
	}
}
