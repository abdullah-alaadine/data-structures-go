package binarysearchtree

type node struct {
	key   int
	left  *node
	right *node
}

type BinarySearchTree struct {
	root *node
}
