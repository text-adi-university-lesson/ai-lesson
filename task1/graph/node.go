package graph

type Node struct {
	name       string
	neighbours []*Node
}

func NewNode(name string) *Node {
	return &Node{name: name}
}

func (n *Node) Value() string {
	return n.name
}

func (n *Node) Neighbours() []*Node {
	return n.neighbours
}

func (n *Node) AddNeighbours(nodes ...*Node) {
	n.neighbours = append(n.neighbours, nodes...)
}
