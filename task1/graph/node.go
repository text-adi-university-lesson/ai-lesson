package graph

type Node struct {
	name string
}

func NewNode(name string) *Node {
	return &Node{name: name}
}

func (n *Node) Value() string {
	return n.name
}
