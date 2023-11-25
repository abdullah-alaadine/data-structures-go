package linkedlist

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head   *node
	length int
}
