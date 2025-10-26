package main

import (
	"ai-lesson/task1/graph"
	"fmt"
)

func main() {
	g := graph.NewGraph()

	fmt.Println("Lабораторна робота - Хмельницький")
	a1 := graph.NewNode("Хмельницький")
	a2 := graph.NewNode("Старокостянтинів")
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
	a16 := graph.NewNode("Маниківці")
	a17 := graph.NewNode("Чотирбоки")
	a18 := graph.NewNode("Наркевичі")
	a19 := graph.NewNode("Кустівці")
	a20 := graph.NewNode("Любар")

	g.AddEdge(graph.NewEdge(a1, a2, 1))
	g.AddEdge(graph.NewEdge(a1, a3, 4))
	g.AddEdge(graph.NewEdge(a1, a4, 5))
	g.AddEdge(graph.NewEdge(a1, a5, 7))
	g.AddEdge(graph.NewEdge(a2, a6, 4))
	g.AddEdge(graph.NewEdge(a4, a6, 4))
	g.AddEdge(graph.NewEdge(a3, a7, 10))
	g.AddEdge(graph.NewEdge(a4, a7, 3))
	g.AddEdge(graph.NewEdge(a4, a8, 6))
	g.AddEdge(graph.NewEdge(a7, a8, 2))
	g.AddEdge(graph.NewEdge(a6, a8, 5))
	g.AddEdge(graph.NewEdge(a1, a9, 7))
	g.AddEdge(graph.NewEdge(a9, a2, 3))
	g.AddEdge(graph.NewEdge(a9, a10, 3))
	g.AddEdge(graph.NewEdge(a2, a10, 1))
	g.AddEdge(graph.NewEdge(a2, a11, 8))
	g.AddEdge(graph.NewEdge(a6, a10, 2))
	g.AddEdge(graph.NewEdge(a1, a13, 9))
	g.AddEdge(graph.NewEdge(a14, a13, 2))
	g.AddEdge(graph.NewEdge(a14, a6, 4))
	g.AddEdge(graph.NewEdge(a15, a6, 8))
	g.AddEdge(graph.NewEdge(a15, a17, 1))
	g.AddEdge(graph.NewEdge(a16, a5, 2))
	g.AddEdge(graph.NewEdge(a18, a1, 3))
	g.AddEdge(graph.NewEdge(a18, a4, 5))
	g.AddEdge(graph.NewEdge(a19, a2, 2))
	g.AddEdge(graph.NewEdge(a19, a17, 9))
	g.AddEdge(graph.NewEdge(a19, a20, 1))
	g.AddEdge(graph.NewEdge(a11, a20, 5))
	g.AddEdge(graph.NewEdge(a12, a7, 3))
	g.AddEdge(graph.NewEdge(a12, a18, 2))
	g.AddEdge(graph.NewEdge(a16, a3, 2))
	g.AddEdge(graph.NewEdge(a12, a13, 4))
	g.AddEdge(graph.NewEdge(a5, a11, 8))
	g.AddEdge(graph.NewEdge(a9, a14, 1))
	g.AddEdge(graph.NewEdge(a1, a14, 2))

	fmt.Println("Count Edges: ", len(g.GetEdges()))

	// Не правильно формується MST. Потрібно зробити тестові дані, та переписати алгоритм MST
	mstGraph := g.GetMST()
	fmt.Println("MST:")
	fmt.Println("Count Edges: ", len(mstGraph.GetEdges()))
	mstGraph.Print()

	//fmt.Println("Neighbours:")
	//for _, node := range mstGraph.GetNodes() {
	//	fmt.Print(node.Value() + ": ")
	//	for _, n := range mstGraph.GetNeighbours(node) {
	//		fmt.Print(n.Value() + ", ")
	//	}
	//	fmt.Println()
	//}
	//mstGraph.RunBFS(a1)
	//mstGraph.RunDFS(a1)
}
