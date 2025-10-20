package graph

type Node struct {
	value      int
	neighbours []*Node
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Neighbours() []*Node {
	return n.neighbours
}

func (n *Node) AddNeighbours(nodes ...*Node) {
	n.neighbours = append(n.neighbours, nodes...)
}
