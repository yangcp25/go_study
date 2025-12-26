package main

import "sort"

func main() {

}

// 128. 最长连续序列
func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, n := range nums {
		set[n] = true
	}

	longest := 0

	for n := range set {
		// 只从“起点”开始
		if !set[n-1] {
			length := 1
			cur := n

			for set[cur+1] {
				cur++
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}

func longestConsecutive2(nums []int) int {
	res := 0    // 用于存储最终的最长连续序列长度
	temMax := 0 // 用于存储当前连续序列的长度
	if len(nums) == 0 {
		return 0 // 如果数组为空，直接返回0
	}
	sort.Ints(nums) // 对数组进行排序，方便后续处理

	// 使用双指针i和j来遍历数组
	for i, j := 0, 0; i < len(nums) && j < len(nums); {
		if nums[j]-nums[i] == 1 { // 如果nums[j]比nums[i]大1，说明是连续序列
			temMax++ // 当前连续序列长度加1
			i++      // 移动i指针
			if temMax > res {
				res = temMax // 更新最长连续序列长度
			}
		} else if nums[j]-nums[i] != 0 { // 如果nums[j]和nums[i]的差值不为0且不为1，说明序列中断
			temMax = 0 // 重置当前连续序列长度
			i = j      // 将i指针移动到j的位置
		} else { // 如果nums[j]和nums[i]相等，说明有重复元素
			i = j // 将i指针移动到j的位置
		}
		j++ // 移动j指针
	}
	return res + 1 // 返回最长连续序列长度，加1是因为初始长度为1
}
