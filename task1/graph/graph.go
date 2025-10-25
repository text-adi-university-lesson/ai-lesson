package graph

import (
	"fmt"
	"sort"
)

type Graph struct {
	edges []*Edge
}

func NewGraph() *Graph {
	return &Graph{}
}

func ImportGraph(edges []*Edge) *Graph {
	return &Graph{edges: edges}
}

func (g *Graph) AddEdge(edge *Edge) {
	g.edges = append(g.edges, edge)
}

func (g *Graph) GetEdges() []*Edge {
	return g.edges
}

func (g *Graph) GetNodes() []*Node {
	nodes := make([]*Node, 0)
	for _, edge := range g.edges {
		n1, n2 := edge.GetNodes()
		if !contains(nodes, n1) {
			nodes = append(nodes, n1)
		}
		if !contains(nodes, n2) {
			nodes = append(nodes, n2)
		}
	}
	return nodes
}

func (g *Graph) GetNeighbours(node *Node) []*Node {
	neighbours := make([]*Node, 0)
	for _, edge := range g.edges {
		n1, n2 := edge.GetNodes()
		if n2 == node && !contains(neighbours, n1) {
			neighbours = append(neighbours, n1)
		}
		if n1 == node && !contains(neighbours, n2) {
			neighbours = append(neighbours, n2)
		}
	}
	return neighbours
}

func (g *Graph) Print() {
	for _, edge := range g.GetEdges() {
		a, b := edge.GetNodes()
		fmt.Println(a.Value(), "-", b.Value())
	}
}

func (g *Graph) getSortedEdge() []*Edge {
	edgesCopy := make([]*Edge, len(g.edges))

	// Копіюємо елементи
	copy(edgesCopy, g.edges)
	sort.Slice(edgesCopy, func(i, j int) bool {
		// відсортувати за вагою ребра
		return edgesCopy[i].GetWight() < edgesCopy[j].GetWight()
	})
	return edgesCopy
}

func (g *Graph) GetMST() *Graph {
	MST := NewGraph()
	for _, edge := range g.getSortedEdge() {
		if !HasCycle(edge) {
			MST.AddEdge(NewEdge(edge.from, edge.to, edge.weight))
		}
	}
	return MST
}
