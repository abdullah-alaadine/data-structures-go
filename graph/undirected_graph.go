package graph

type Node struct {
	Neighbors []*Node
}

type Graph struct {
	nodes map[int]*Node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]*Node),
	}
}
