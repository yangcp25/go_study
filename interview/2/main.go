package main

import "fmt"

func main() {
	//
	tree := make([][]node, 4)
	tree[1] = []node{node{2, 1}}
	tree[1] = []node{node{3, 2}}
	tree[2] = []node{node{4, 5}}
	tree[2] = []node{node{5, 3}}

	path := make(map[int]int)
	for _, v := range tree {
		for _, n := range v {
			path[n.key] = Max(path[n.key], n.dist)
		}
	}
	fmt.Println(path)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type node struct {
	key  int
	dist int
}
