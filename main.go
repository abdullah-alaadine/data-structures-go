package main

import (
	"github.com/knbr13/data-structures-go/graph"
)

func main() {
	g := graph.NewGraph()
	n1 := g.AddNode(10)
	n2 := g.AddNode(11)
	n3 := g.AddNode(12)
	n4 := g.AddNode(13)
	g.AddEdge(n1.ID(), n2.ID())
	g.AddEdge(n2.ID(), n3.ID())
	g.AddEdge(n3.ID(), n4.ID())
	g.AddEdge(n4.ID(), n1.ID())
	g.PrintGraph()
}
