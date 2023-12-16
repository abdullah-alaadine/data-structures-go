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

func (g *Graph) AddEdge(firstNodeID, secondNodeID int) {
	firstNode := g.nodes[firstNodeID]
	secondNode := g.nodes[secondNodeID]

	firstNode.Neighbors = append(firstNode.Neighbors, secondNode)
	secondNode.Neighbors = append(secondNode.Neighbors, firstNode)
}

func (g *Graph) RemoveNode(nodeID int) *Node {
	node, ok := g.nodes[nodeID]
	if !ok {
		return nil
	}

	for _, node := range node.Neighbors {
		for i := range g.nodes[node.ID].Neighbors {
			if g.nodes[node.ID].Neighbors[i].ID == node.ID {
				g.nodes[node.ID].Neighbors = append(g.nodes[node.ID].Neighbors[:i], g.nodes[node.ID].Neighbors[i+1:]...)
			}
		}
	}

	delete(g.nodes, nodeID)
	return node
}
