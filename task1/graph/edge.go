package graph

type Edge struct {
	from, to *Node
	weight   int
}

func (e *Edge) GetWight() int {
	return e.weight
}

func (e *Edge) GetNodes() (*Node, *Node) {
	return e.from, e.to
}

func NewEdge(from *Node, to *Node, weight int) *Edge {
	return &Edge{from: from, to: to, weight: weight}
}
