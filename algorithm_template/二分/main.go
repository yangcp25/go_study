package main

func main() {

}

//

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func BinarySearchIndex(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		// 左半边有序
		if nums[l] <= nums[mid] {
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // 右半边有序
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

// 162 寻找峰值
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[mid+1] {
			left = mid + 1 // 上坡，往右
		} else {
			right = mid // 下坡，往左
		}
	}
	return left
}

// 34 在排序数组中查找元素范围
func searchRange(nums []int, target int) []int {
	left := findPosFirst(nums, target)
	right := findPosLast(nums, target)
	return []int{left, right}
}
func findPosFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	res := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			res = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}

func findPosLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	res := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			res = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
