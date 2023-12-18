package main

import (
	"fmt"
	"sort"
)

func main() {
	testArray := []int{3, 1, 2, 6, 4, 5, 7, 8, 10}
	sort.Ints(testArray)
	fmt.Printf("数组为%+v \n", testArray)

	searchData := 9
	check := binarySearchV2(testArray, 9, 0, len(testArray)-1)

	fmt.Printf("查找%d结果%d \n", searchData, check)

	searchData = 6
	check2 := binarySearchV2(testArray, searchData, 0, len(testArray)-1)

	fmt.Printf("查找%d结果%d \n", searchData, check2)

	searchData = 8
	testArray = []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}
	check3 := binarySearchV3(testArray, searchData, 0, len(testArray)-1)
	check4 := binarySearchV3_1(testArray, searchData, 0, len(testArray)-1)

	fmt.Printf("查找%d结果%d \n", searchData, check3)
	fmt.Printf("查找%d结果%d \n", searchData, check4)

	testArray = []int{3, 4, 5, 5, 10}
	searchData = 5
	check5 := binarySearchV5(testArray, searchData, 0, len(testArray)-1)
	fmt.Printf("查找%d结果%d \n", searchData, check5)

	testArray = []int{3, 5, 6, 8, 9, 10}
	searchData = 6
	check6 := binarySearchV7(testArray, searchData, 0, len(testArray)-1)
	fmt.Printf("查找%d结果%d \n", searchData, check6)

}

// 二分查找函数递归
func binarySearch[T int | string](data []T, searchData T, start, end int) int {

	if start > end {
		return -1
	}

	mid := int((end + start) / 2)

	if data[mid] == searchData {
		return mid
	} else if data[mid] > searchData {
		return binarySearch(data, searchData, start, mid-1)
	} else {
		return binarySearch(data, searchData, mid+1, end)
	}
}

// 二分查找函数非递归写法
func binarySearchV2[T int | string](data []T, searchData T, start, end int) int {
	for start <= end {
		currenKey := int((start + end) / 2)
		if data[currenKey] == searchData {
			return currenKey
		} else if data[currenKey] < searchData {
			start = currenKey + 1
		} else {
			end = currenKey - 1
		}
	}
	return -1
}

// 变体一：查找第一个值等于给定值的元素
/**
(1) 当找到第一个数据符合条件的数据 就往左边找
*/
func binarySearchV3[T int | string](data []T, searchData T, start, end int) int {
	findKey := -1
	for start <= end {
		currenKey := int((start + end) / 2)
		if data[currenKey] == searchData {
			findKey = currenKey
			break
		} else if data[currenKey] < searchData {
			start = currenKey + 1
		} else {
			end = currenKey - 1
		}
	}
	if findKey == -1 {
		return -1
	}
	// 这种虽然用了for 但是一般循环不了几次 除非重复数据特别多的情况
	for findKey > 0 {
		if data[findKey-1] != searchData {
			return findKey
		}
		findKey--
	}
	return findKey
}

// 上面的改进
func binarySearchV3_1[T int | string](data []T, searchData T, start, end int) int {
	for start <= end {
		currenKey := int((start + end) / 2)
		if data[currenKey] == searchData {
			if data[currenKey-1] != searchData {
				return currenKey
			}
			end = currenKey - 1
		} else if data[currenKey] < searchData {
			start = currenKey + 1
		} else {
			end = currenKey - 1
		}
	}
	return -1
}

// 变体二：查找最后一个值等于给定值的元素
func binarySearchV4[T int | string](data []T, searchData T, start, end int) int {
	for start <= end {
		currenKey := int((start + end) / 2)
		if data[currenKey] == searchData {
			if data[currenKey+1] != searchData {
				return currenKey
			}
			// 相等的情况 中间值该往哪边走？ 这里往右/高的数据走
			start = currenKey + 1
		} else if data[currenKey] < searchData {
			start = currenKey + 1
		} else {
			end = currenKey - 1
		}
	}
	return -1
}

// 变体三：查找第一个大于等于给定值的元素
func binarySearchV5[T int | string](data []T, searchData T, start, end int) int {
	for start <= end {
		currenKey := int((start + end) / 2)
		if data[currenKey] >= searchData {
			if currenKey == 0 || data[currenKey-1] < searchData {
				return currenKey
			}
			end = currenKey - 1
		} else {
			start = currenKey + 1
		}
	}
	return -1
}

// 变体三：查找第一个大于给定值的元素
func binarySearchV6[T int | string](data []T, searchData T, start, end int) int {
	for start <= end {
		currenKey := int((start + end) / 2)
		if data[currenKey] > searchData {
			if currenKey == 0 || data[currenKey-1] < searchData || data[currenKey-1] == searchData {
				return currenKey
			}
			end = currenKey - 1
		} else {
			start = currenKey + 1
		}
	}
	return -1
}

// 查找最后一个小于等于给定值的元素
func binarySearchV7[T int | string](data []T, searchData T, start, end int) int {
	for start <= end {
		currenKey := int((start + end) / 2)
		if data[currenKey] <= searchData {
			if currenKey == end || data[currenKey+1] > searchData {
				return currenKey
			}
			start = currenKey + 1
		} else {
			end = currenKey - 1
		}
	}
	return -1
}
