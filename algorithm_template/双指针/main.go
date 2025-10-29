package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{-1, 0}
	//target := -1
	//res := towSum(nums, target)
	//fmt.Println(res)

	nums := []int{-1, 0, 1, 2, -1, -4}
	//fmt.Println(threeSum(nums))

	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

// 两数之和
func towSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, number := range numbers {
		if j, ok := m[target-number]; ok {
			return []int{j, i}
		}
		m[number] = i
	}
	return nil
}

// 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

// 盛雨水
func maxArea(height []int) int {
	left, right, res := 0, len(height)-1, 0
	for left < right {
		area := (right - left) * Min(height[left], height[right])
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		if area > res {
			res = area
		}
	}
	return res
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 快速排序
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	// 分区操作
	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

// 98 验证二叉搜索树
