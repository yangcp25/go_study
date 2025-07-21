package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 0
	//   1 -> 4 -> 5
	// 0 2 -> 6 -> 7
	//   3 ->^
	g := make([][]int, 8)
	//g = append(g, []int{1, 2, 3}
	g[0] = []int{1, 2, 3}
	g[1] = []int{4}
	g[2] = []int{6}
	g[3] = []int{6}
	g[4] = []int{5}
	g[6] = []int{7}

	g2 := make(map[int][]int)
	g2[0] = []int{1, 2, 3}
	g2[1] = []int{4}
	g2[2] = []int{6}
	g2[3] = []int{6}
	g2[4] = []int{5}
	g2[6] = []int{7}

	graphDFS(0, g)
	//graphBFS(g)

	fmt.Println("-----")

	graphBFSWithQueue(g2)
	fmt.Println("-----")
	listC := graphBFSWithC(g2)
	fmt.Println(listC)
}
func graphDFS(root int, graph [][]int) {
	fmt.Println(root)
	if graph[root] == nil {
		return
	}
	//fmt.Println(root)

	for i := 0; i < len(graph[root]); i++ {
		graphDFS(graph[root][i], graph)
	}
	return
}

func graphBFSWithQueue(graph map[int][]int) {
	queue := make([]int, 0)
	queue = append(queue, 0)
	fmt.Println(0)
	visited := make(map[int]bool)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		g := graph[node]
		size := len(g)
		for i := 0; i < size; i++ {
			if visited[g[i]] {
				continue
			} else {
				fmt.Println(g[i])
				queue = append(queue, g[i])
				visited[g[i]] = true
			}
		}
	}
}
func graphBFSWithC(graph map[int][]int) []int {
	queue := list.New()
	order := make([]int, 0)
	visited := make(map[int]bool)
	queue.PushBack(0)

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(int)

		order = append(order, node)

		neighbor := graph[node]
		size := len(neighbor)
		for i := 0; i < size; i++ {
			if visited[neighbor[i]] {
				continue
			} else {
				queue.PushBack(neighbor[i])
				visited[neighbor[i]] = true
			}
		}
	}
	return order
}
