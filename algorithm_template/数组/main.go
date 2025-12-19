package main

import "sort"

func main() {

}

// 88 合并2个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

// 1.两数之和
// 15.三数之和
// 2.两数相加
// 569 和为k的子数组
// 53 最大子数组和
// 56 合并区间
func merge2(intervals [][]int) [][]int {
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
