package main

import (
	heap2 "container/heap"
	"fmt"
)

func main() {
	res := findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1})
	fmt.Println(res)
}

// 。k = 2, w = 0, profit = [1,2,3], capital = [0,1,1]
// k = 3, w = 0, profit = [1,2,3], capital = [0,1,2]
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	// sum := 0
	// 预处理数据，计算每个项目的利润是多少
	// 将所有项目放入最小堆，每次弹出一个，当满足 <= 当前资本，将项目加入最大堆
	// 如果最大堆的数量为0，跳出循环，返回结果
	// 否则从最大堆取出数据，加入利润，并标记已使用
	// 每次从最小堆弹出数据，需要判断是否已经使用过了

	projects := make([]project, 0)
	for key, val := range capital {
		projects = append(projects, project{
			key:         key,
			profit:      val,
			pureProfits: profits[key],
		})
	}

	//
	check := make(map[int]bool)

	for len(check) < k {
		heapMax := &MaxHeap{}
		heap2.Init(heapMax)
		for i := 0; i < k; i++ {
			if capital[i] <= w && !check[i] {
				heap2.Push(heapMax, project{
					key:         i,
					pureProfits: profits[i] - capital[i],
					profit:      profits[i],
				})
			}
		}
		if heapMax.Len() > 0 {
			project := heap2.Pop(heapMax).(project)
			w += project.pureProfits
			check[project.key] = true
		} else {
			break
		}
	}

	return w
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type project struct {
	key         int
	pureProfits int
	profit      int
}
type MaxHeap []project

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i].pureProfits > m[j].pureProfits
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(project))
}

func (m *MaxHeap) Pop() any {
	old := *m
	n := len(old)
	data := old[n-1]
	*m = old[:n-1]
	return data
}
