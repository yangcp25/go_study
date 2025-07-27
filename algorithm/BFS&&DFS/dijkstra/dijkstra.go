package main

func main() {
	// 0
	//   1 -> 4 -> 5
	// 0 2 -> 6 -> 7
	//   3 ->^

	graph := make([][]Edge, 8)
	graph[0] = []Edge{
		{1, 1},
		{2, 1},
		{3, 1},
	}
	graph[1] = []Edge{
		{4, 1},
	}
	graph[2] = []Edge{
		{6, 1},
	}
	graph[3] = []Edge{
		{6, 1},
	}
	graph[4] = []Edge{
		{5, 1},
	}
	graph[5] = []Edge{}
	graph[7] = []Edge{}
	graph[6] = []Edge{
		{7, 1},
	}
}

type Edge struct {
	To     int
	Weight int
}
