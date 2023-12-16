package graph

import "fmt"

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

func (g *Graph) RemoveEdge(firstNodeID, secondNodeID int) {
	firstNode := g.nodes[firstNodeID]
	secondNode := g.nodes[secondNodeID]

	for i := range firstNode.Neighbors {
		if firstNode.Neighbors[i].ID == secondNodeID {
			firstNode.Neighbors = append(firstNode.Neighbors[:i], firstNode.Neighbors[i+1:]...)
		}
	}
	for i := range secondNode.Neighbors {
		if secondNode.Neighbors[i].ID == firstNodeID {
			secondNode.Neighbors = append(secondNode.Neighbors[:i], secondNode.Neighbors[i+1:]...)
		}
	}
}

func (g *Graph) PrintGraph() {
	for _, node := range g.nodes {
		fmt.Printf("Node %d: ", node.ID)
		for _, neighbor := range node.Neighbors {
			fmt.Printf("%d - ", neighbor.ID)
		}
		fmt.Print("\n")
	}
}
