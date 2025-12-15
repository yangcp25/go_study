package main

import "sort"

func main() {

}

//56. 合并区间（Merge Intervals）

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 1. 按起点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)

	// 2. 当前区间
	curStart := intervals[0][0]
	curEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		s, e := intervals[i][0], intervals[i][1]

		if s <= curEnd {
			// 有重叠，合并
			curEnd = max(curEnd, e)
		} else {
			// 无重叠，保存当前区间
			res = append(res, []int{curStart, curEnd})
			curStart = s
			curEnd = e
		}
	}

	// 3. 别忘了最后一个区间
	res = append(res, []int{curStart, curEnd})

	return res
}
