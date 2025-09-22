package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
*

代码
测试用例
测试结果
测试结果
502. IPO
已解答
困难
相关标签
premium lock icon
相关企业
假设 力扣（LeetCode）即将开始 IPO 。为了以更高的价格将股票卖给风险投资公司，力扣 希望在 IPO 之前开展一些项目以增加其资本。 由于资源有限，它只能在 IPO 之前完成最多 k 个不同的项目。帮助 力扣 设计完成最多 k 个不同项目后得到最大总资本的方式。

给你 n 个项目。对于每个项目 i ，它都有一个纯利润 profits[i] ，和启动该项目需要的最小资本 capital[i] 。

最初，你的资本为 w 。当你完成一个项目时，你将获得纯利润，且利润将被添加到你的总资本中。

总而言之，从给定项目中选择 最多 k 个不同项目的列表，以 最大化最终资本 ，并输出最终可获得的最多资本。

答案保证在 32 位有符号整数范围内。

示例 1：

输入：k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
输出：4
解释：
由于你的初始资本为 0，你仅可以从 0 号项目开始。
在完成后，你将获得 1 的利润，你的总资本将变为 1。
此时你可以选择开始 1 号或 2 号项目。
由于你最多可以选择两个项目，所以你需要完成 2 号项目以获得最大的资本。
因此，输出最后最大化的资本，为 0 + 1 + 3 = 4。
示例 2：

输入：k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
输出：6
*/
func main() {
	res := findMaximizedCapital(10, 0, []int{1, 2, 3}, []int{0, 1, 2})
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
			capital: val,
			profit:  profits[key],
		})
	}
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	idx := 0
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for i := 0; i < k; i++ {
		// 从project 拿到可以的资本列表 ，然后选择最大的利润
		j := idx
		for ; j < len(projects); j++ {
			if projects[j].capital <= w {
				heap.Push(maxHeap, projects[j])
				idx = j
			}
		}
		if maxHeap.Len() == 0 {
			break
		}
		cur := heap.Pop(maxHeap).(project)
		w += cur.profit
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
	capital int
	profit  int
}
type MaxHeap []project

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i].profit > m[j].profit
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
