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

func Task2() {
	fmt.Println("Lабораторна робота - Хмельницький")
	a1 := graph.NewNode("Хмельницький")
	a2 := graph.NewNode("Старокоснятинів")
	a3 := graph.NewNode("Ярмолинці")
	a4 := graph.NewNode("Волочиськ")
	a5 := graph.NewNode("Летичів")
	a6 := graph.NewNode("Теофіполь")
	a7 := graph.NewNode("Сатанів")
	a8 := graph.NewNode("Тернопіль")
	a9 := graph.NewNode("Красилів")
	a10 := graph.NewNode("Антоніни")
	a11 := graph.NewNode("Адампіль")
	a12 := graph.NewNode("Війтівці")
	a13 := graph.NewNode("Купіль")
	a14 := graph.NewNode("Базалія")
	a15 := graph.NewNode("Білогір'я")
	a16 := graph.NewNode("Білогородка")
	a17 := graph.NewNode("Чотирбоки")
	a18 := graph.NewNode("Наркевичі")
	a19 := graph.NewNode("Кустівці")
	a20 := graph.NewNode("Любар")

	edges := []*graph.Edge{
		// якщо потрібно одразу ініціалізувати масив
		graph.NewEdge(a1, a2, 1),
		graph.NewEdge(a1, a3, 4),
		graph.NewEdge(a1, a4, 5),
		graph.NewEdge(a1, a5, 7),
		graph.NewEdge(a2, a6, 4),
		graph.NewEdge(a2, a6, 1),
		graph.NewEdge(a4, a6, 4),
		graph.NewEdge(a3, a7, 10),
		graph.NewEdge(a4, a7, 3),
		graph.NewEdge(a4, a8, 6),
		graph.NewEdge(a7, a8, 2),
		graph.NewEdge(a6, a8, 5),
		graph.NewEdge(a1, a9, 7),
		graph.NewEdge(a9, a2, 3),
		graph.NewEdge(a9, a10, 3),
		graph.NewEdge(a2, a10, 1),
		graph.NewEdge(a2, a10, 8),
		graph.NewEdge(a6, a10, 2),
		graph.NewEdge(a1, a13, 9),
		graph.NewEdge(a14, a13, 2),
		graph.NewEdge(a14, a6, 4),
		graph.NewEdge(a15, a6, 8),
		graph.NewEdge(a15, a16, 1),
		graph.NewEdge(a17, a16, 2),
		graph.NewEdge(a17, a16, 7),
		graph.NewEdge(a18, a1, 3),
		graph.NewEdge(a18, a4, 5),
		graph.NewEdge(a19, a2, 2),
		graph.NewEdge(a19, a17, 9),
		graph.NewEdge(a19, a20, 1),
		graph.NewEdge(a11, a20, 5),
		graph.NewEdge(a11, a20, 6),
		graph.NewEdge(a12, a20, 3),
		graph.NewEdge(a12, a11, 2),
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

	fmt.Println("MST:")
	for _, edge := range MST {
		a, b := edge.GetNodes()
		fmt.Println(a.Value(), "-", b.Value())
	}
}

func main() {
	Task2()

	checkBFS()
	checkDFS()

}
