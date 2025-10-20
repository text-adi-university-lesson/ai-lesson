package main

import (
	"ai-lesson/task1/graph"
	"fmt"
)

func contains(slice []*graph.Node, value *graph.Node) bool {
	for _, v := range slice {
		if v.Value() == value.Value() {
			return true
		}
	}
	return false
}

func BFS(n *graph.Node) {
	visited := make([]*graph.Node, 0)
	queue := make([]*graph.Node, 0)

	visited = append(visited, n)
	queue = append(queue, n)

	for len(queue) > 0 {
		currentNode := queue[0]

		fmt.Println("Поточний Node", currentNode.Value(), " ", "Поточні відвідувані елементи", visited, " ", "Поточна черга", queue)

		// Додати сусідів до черги
		for _, neighbour := range currentNode.Neighbours() {
			if !contains(visited, neighbour) {
				queue = append(queue, neighbour)
				visited = append(visited, neighbour)

			}
		}
		queue = queue[1:] // Видалити перший елемент

	}
}

func DFS(n *graph.Node) {
	visited := make([]*graph.Node, 0)
	stack := make([]*graph.Node, 0)

	currentNode := n

	for {
		visited = append(visited, currentNode)

		// Додати сусідів до стеку
		for _, neighbour := range currentNode.Neighbours() {
			if !contains(visited, neighbour) {
				stack = append(stack, neighbour)
			}
		}
		fmt.Println("Поточний Node", currentNode.Value(), " ", "Поточні відвідувані елементи", visited, " ", "Поточний стек", stack)

		if len(stack) == 0 {
			// вийти з програми, якщо стек порожній
			break
		}
		currentNode = stack[0]
		stack = stack[1:]

	}
}

func checkBFS() {
	fmt.Println("Метод пошуку в ширину (BFS):")
	a := graph.NewNode(0)
	b := graph.NewNode(1)
	c := graph.NewNode(2)
	d := graph.NewNode(3)
	e := graph.NewNode(4)

	a.AddNeighbours(b, c)
	b.AddNeighbours(a, c, d)
	c.AddNeighbours(a, b, e)
	d.AddNeighbours(b, e)
	e.AddNeighbours(c, d)

	BFS(a)
}

func checkDFS() {
	fmt.Println("Метод пошуку в ширину (DFS):")
	a := graph.NewNode(0)
	b := graph.NewNode(1)
	c := graph.NewNode(2)
	d := graph.NewNode(3)
	e := graph.NewNode(4)

	a.AddNeighbours(b, c, d)
	b.AddNeighbours(a)
	c.AddNeighbours(a, d, e)
	d.AddNeighbours(a, c)
	e.AddNeighbours(c)

	DFS(a)
}

func main() {
	checkBFS()
	checkDFS()

}
