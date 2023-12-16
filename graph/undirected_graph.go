package graph

import "fmt"

var id int = 0

type node struct {
	id        int
	Value     any
	neighbors []*node
}

func (n *node) ID() int { return n.id }

type Graph struct {
	nodes map[int]*node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]*node),
	}
}

func newNode(value any) *node {
	id++
	return &node{
		id:    id,
		Value: value,
	}
}

func (g *Graph) AddNode(value any) {
	n := newNode(value)
	g.nodes[n.id] = n
}

func (g *Graph) AddEdge(firstNodeID, secondNodeID int) {
	firstNode := g.nodes[firstNodeID]
	secondNode := g.nodes[secondNodeID]

	firstNode.neighbors = append(firstNode.neighbors, secondNode)
	secondNode.neighbors = append(secondNode.neighbors, firstNode)
}

func (g *Graph) RemoveNode(nodeID int) *node {
	node, ok := g.nodes[nodeID]
	if !ok {
		return nil
	}

	for _, node := range node.neighbors {
		for i := range g.nodes[node.id].neighbors {
			if g.nodes[node.id].neighbors[i].id == node.id {
				g.nodes[node.id].neighbors = append(g.nodes[node.id].neighbors[:i], g.nodes[node.id].neighbors[i+1:]...)
			}
		}
	}

	delete(g.nodes, nodeID)
	return node
}

func (g *Graph) RemoveEdge(firstNodeID, secondNodeID int) {
	firstNode := g.nodes[firstNodeID]
	secondNode := g.nodes[secondNodeID]

	for i := range firstNode.neighbors {
		if firstNode.neighbors[i].id == secondNodeID {
			firstNode.neighbors = append(firstNode.neighbors[:i], firstNode.neighbors[i+1:]...)
		}
	}
	for i := range secondNode.neighbors {
		if secondNode.neighbors[i].id == firstNodeID {
			secondNode.neighbors = append(secondNode.neighbors[:i], secondNode.neighbors[i+1:]...)
		}
	}
}

func (g *Graph) PrintGraph() {
	for _, node := range g.nodes {
		fmt.Printf("Node %d: ", node.id)
		for _, neighbor := range node.neighbors {
			fmt.Printf("%d - ", neighbor.id)
		}
		fmt.Print("\n")
	}
}
