package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	/**
	5 7
	1 4
	2 1
	2 3
	2 4
	3 4
	3 5
	4 5
	2
	*/
	n := 0
	t := 0

	fmt.Scan(&n, &t)

	g := NewGraph(n + 1)
	fmt.Println(g)
	for i := 0; i < t; i++ {
		u, v := 0, 0
		fmt.Scan(&u, &v)
		g.AddEdge(u, v)
	}

	start := 0
	fmt.Scan(&start)

	distance := Djikstra(g, start)

	sort.Ints(distance)
	maxL := distance[len(distance)-1]

	fmt.Println(maxL * 2)

}

func Djikstra(g *Graph, start int) []int {
	dist := make([]int, g.n+1)
	for k, _ := range dist {
		dist[k] = -1
	}
	dist[start] = 0

	// 初始化最小堆
	priorityQueue := make(PriorityQueue, 0)
	heap.Init(&priorityQueue)
	heap.Push(&priorityQueue, &Item{
		Vertex:   start,
		Distance: 0,
	})

	for len(priorityQueue) > 0 {
		item := heap.Pop(&priorityQueue).(*Item)
		distance := item.Distance
		node := item.Vertex

		if dist[node] != -1 && distance > dist[node] {
			continue
		}
		for _, neighbor := range g.adj[node] {
			if dist[neighbor] == -1 || dist[node]+1 < dist[neighbor] {
				dist[neighbor] = dist[node] + 1
				heap.Push(&priorityQueue, &Item{
					Vertex:   neighbor,
					Distance: dist[node] + 1,
				})
			}
		}
	}

	return dist
}

// 定义图
type Graph struct {
	// 领接表
	adj [][]int
	n   int
}

func NewGraph(n int) *Graph {
	adj := make([][]int, n)
	for key := range adj {
		adj[key] = make([]int, 0)
	}
	return &Graph{
		adj: adj,
		n:   n,
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u) // 无向图
}

// 初始化最小堆（优先级队列）

type Item struct {
	Vertex   int
	Distance int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = j, i
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[0]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
