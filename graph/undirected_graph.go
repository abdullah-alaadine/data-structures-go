package graph

type Node struct {
	ID        int
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

func (g *Graph) AddNode(n *Node) {
	g.nodes[n.ID] = n
}
