package graph

type Node struct {
	Neighbors []*Node
}

type Graph struct {
	nodes map[int]*Node
}
