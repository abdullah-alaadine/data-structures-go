package doublylinkedlist

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
