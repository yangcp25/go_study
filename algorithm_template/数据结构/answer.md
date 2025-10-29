``` go
// 链表
type LinkNode struct {
	Val  int
	Next *LinkNode
}

// 双向链表
type DLinkNode struct {
	Val  int
	Pre  *DLinkNode
	Next *DLinkNode
}

// 栈
type Stack struct {
	data []int
}

// 队列
type Queue struct {
	data []int
}

// 树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// N叉树
type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

// 图
type GraphNode struct {
	Val      int
	Neighbor []*GraphNode
}

// 图（邻接表（推荐方式））

type Graph struct {
    edges map[string][]string
}

func NewGraph() *Graph {
    return &Graph{edges: make(map[string][]string)}
}

func (g *Graph) AddEdge(from, to string) {
    g.edges[from] = append(g.edges[from], to)
}

// 图（带权图（Weighted Graph））

type Graph struct {
    edges map[string]map[string]int
}

func NewGraph() *Graph {
    return &Graph{edges: make(map[string]map[string]int)}
}

func (g *Graph) AddEdge(from, to string, weight int) {
    if g.edges[from] == nil {
        g.edges[from] = make(map[string]int)
    }
    g.edges[from][to] = weight
}

// 图的遍历（递归）

func (g *Graph) DFS(start string, visited map[string]bool) {
    if visited[start] {
        return
    }
    visited[start] = true
    fmt.Println("visit:", start)

    for _, neighbor := range g.edges[start] {
        g.DFS(neighbor, visited)
    }
}

// BFS 广度优先搜索（队列）
func (g *Graph) BFS(start string) {
    visited := make(map[string]bool)
    queue := []string{start}
    visited[start] = true

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        fmt.Println("visit:", node)

        for _, neighbor := range g.edges[node] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
}


// 堆/优先级队列

type IntHeapItem struct {
	val int
}

type IntHeap []IntHeapItem



// 矩阵 matrix
//matrix := [][]int {
//	{1, 2 ,3},
//	{1, 2 ,3},
//}

```