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
