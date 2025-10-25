package graph

var components = make([][]*Node, 0)

func contains(slice []*Node, value *Node) bool {
	for _, v := range slice {
		if v.Value() == value.Value() {
			return true
		}
	}
	return false
}

func HasCycle(edge *Edge) bool {
	n1, n2 := edge.GetNodes()
	existsInComponent := false
	for _, component := range components {
		if contains(component, n1) {
			existsInComponent = true
			break
		}
	}
	if !existsInComponent {
		components = append(components, []*Node{n1})
	}

	for i, component := range components {
		if contains(component, n1) {
			if contains(component, n2) {
				return true
			} else {
				components[i] = append(component, n2)
				return false
			}
		}
	}
	return false
}
