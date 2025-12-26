package main

import (
	"math"
	"sort"
)

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
// 128. 最长连续序列
// 300. 最长递增子序列
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

//88. 合并2个有序数组

func mergeTwoArray(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

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

// 4. 寻找正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 保证 nums1 是更短的数组
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	total := m + n
	half := total / 2

	left, right := 0, m

	for left <= right {
		i := (left + right) / 2 // nums1 切点
		j := half - i           // nums2 切点

		// 处理边界：越界时用 ±∞ 代替
		Aleft := math.Inf(-1)
		if i > 0 {
			Aleft = float64(nums1[i-1])
		}
		Aright := math.Inf(1)
		if i < m {
			Aright = float64(nums1[i])
		}

		Bleft := math.Inf(-1)
		if j > 0 {
			Bleft = float64(nums2[j-1])
		}
		Bright := math.Inf(1)
		if j < n {
			Bright = float64(nums2[j])
		}

		// 是否满足合法划分
		if Aleft <= Bright && Bleft <= Aright {
			// 已找到正确划分
			if total%2 == 1 {
				return math.Min(Aright, Bright)
			}
			return (math.Max(Aleft, Bleft) + math.Min(Aright, Bright)) / 2
		}

		// 移动二分边界
		if Aleft > Bright {
			right = i - 1
		} else {
			left = i + 1
		}
	}

	return 0
}
