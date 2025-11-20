package main

import (
	"container/heap"
	"math"
)

func main() {

}

type Edge struct {
	to, weight int
}

type Node struct {
	id, dist int
}

type MinPQ []Node

func (pq MinPQ) Len() int           { return len(pq) }
func (pq MinPQ) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq MinPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *MinPQ) Push(x interface{}) {
	*pq = append(*pq, x.(Node))
}

func (pq *MinPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[:n-1]
	return node
}

// 123
// -1 2
func Dijkstra(n int, graph [][]Edge, start int) []int {
	path := make([]int, n)
	for i := range path {
		path[i] = math.MaxInt32
	}
	path[start] = 0
	pq := &MinPQ{}
	heap.Init(pq)
	heap.Push(pq, Node{id: start, dist: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Node)
		if cur.dist > path[cur.id] {
			continue
		}
		for _, e := range graph[cur.id] {
			if path[cur.id]+e.weight < path[e.to] {
				path[e.to] = path[cur.id] + e.weight
				heap.Push(pq, Node{id: e.to, dist: path[e.to]})
			}
		}
	}
	return path
}
