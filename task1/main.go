package main

import (
	"ai-lesson/task1/graph"
	"fmt"
	"sort"
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
	a := graph.NewNode("0")
	b := graph.NewNode("1")
	c := graph.NewNode("2")
	d := graph.NewNode("3")
	e := graph.NewNode("4")

	a.AddNeighbours(b, c)
	b.AddNeighbours(a, c, d)
	c.AddNeighbours(a, b, e)
	d.AddNeighbours(b, e)
	e.AddNeighbours(c, d)

	BFS(a)
}

func checkDFS() {
	fmt.Println("Метод пошуку в ширину (DFS):")
	a := graph.NewNode("0")
	b := graph.NewNode("1")
	c := graph.NewNode("2")
	d := graph.NewNode("3")
	e := graph.NewNode("4")

	a.AddNeighbours(b, c, d)
	b.AddNeighbours(a)
	c.AddNeighbours(a, d, e)
	d.AddNeighbours(a, c)
	e.AddNeighbours(c)

	DFS(a)
}

func Lar1() {
	fmt.Println("Lабораторна робота - Хмельницький")
	a1 := graph.NewNode("Хмельницький")
	a2 := graph.NewNode("Старокоснятинів")
	a3 := graph.NewNode("Голосків")
	a4 := graph.NewNode("Меджибіж")
	a5 := graph.NewNode("Летичів")
	a6 := graph.NewNode("Стара Синява")
	a7 := graph.NewNode("Адампіль")
	a8 := graph.NewNode("Левківка")
	a9 := graph.NewNode("Красилів")
	a10 := graph.NewNode("Антоніни")
	a11 := graph.NewNode("Теофіполь")
	a12 := graph.NewNode("Війтівці")
	a13 := graph.NewNode("Купіль")
	a14 := graph.NewNode("Базалія")
	a15 := graph.NewNode("Білогір'я")
	a16 := graph.NewNode("Білогородка")
	a17 := graph.NewNode("Чотирбоки")
	a18 := graph.NewNode("Гринців")
	a19 := graph.NewNode("Кустівці")
	a20 := graph.NewNode("Любар")

	_ = []*graph.Node{a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20}
	edges := []*graph.Edge{
		// якщо потрібно одразу ініціалізувати масив
		graph.NewEdge(a1, a2, 1),
		graph.NewEdge(a2, a3, 2),
		graph.NewEdge(a3, a4, 3),
		graph.NewEdge(a1, a4, 4),
		graph.NewEdge(a2, a4, 5),
	}

	sort.Slice(edges, func(i, j int) bool {
		// відсортувати за вагою ребра
		return edges[i].GetWight() < edges[j].GetWight()
	})

	MST := make([]*graph.Edge, 0)
	for _, edge := range edges {
		if !graph.HasCycle(edge) {
			MST = append(MST, edge)
		}
	}

	fmt.Println("Edges:", edges)
}

func main() {
	Lar1()

	checkBFS()
	checkDFS()

}
