package binarysearchtree

type node struct {
	key   int
	left  *node
	right *node
}

type BinarySearchTree struct {
	root *node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (t *BinarySearchTree) Insert(key int) {
	newNode := &node{key: key}
	if t.root == nil {
		t.root = newNode
		return
	}
	insertHelper(t.root, newNode)
}

func insertHelper(root *node, node *node) {
	if root.key < node.key {
		if root.right == nil {
			root.right = node
			return
		}
		insertHelper(root.right, node)
	}
	if root.key > node.key {
		if root.left == nil {
			root.left = node
			return
		}
		insertHelper(root.left, node)
	}
}
