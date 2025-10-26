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
func existsGroup(node *Node) bool {
	for _, component := range components {
		if contains(component, node) {
			return true
		}
		continue
	}
	return false
}
func getIDGroup(node *Node) int {
	for i, component := range components {
		if contains(component, node) {
			return i
		}
	}
	return -1
}

func HasCycle(edge *Edge) bool {
	n1, n2 := edge.GetNodes()

	existsInGroupN1 := existsGroup(n1)
	existsInGroupN2 := existsGroup(n2)

	if !existsInGroupN1 && !existsInGroupN2 {
		components = append(components, []*Node{n1, n2})
	}

	if existsInGroupN1 && existsInGroupN2 {
		groupID1 := getIDGroup(n1)
		groupID2 := getIDGroup(n2)

		if groupID1 == groupID2 {
			return true
		} else {
			//	потрібно node з однієї групи об'єднати в ноди іншої групи
			components[groupID1] = append(components[groupID1], components[groupID2]...)
			components = append(components[:groupID2], components[groupID2+1:]...)
		}

	}

	if existsInGroupN1 && !existsInGroupN2 {
		groupID := getIDGroup(n1)
		if groupID == -1 {
			panic("Element is not a group")
		}
		components[groupID] = append(components[groupID], n2)
	}

	if existsInGroupN2 && !existsInGroupN1 {
		groupID := getIDGroup(n2)
		if groupID == -1 {
			panic("Element is not a group")
		}
		components[groupID] = append(components[groupID], n1)
	}

	return false
}
