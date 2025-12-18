package main

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
