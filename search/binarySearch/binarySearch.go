package main

import (
	"fmt"
	"sort"
)

// 查找算法
// 二分查找（）
func main() {
	intArray := []int{4, 6, 5, 3, 1, 8, 2, 7}
	sort.Ints(intArray)
	res := binarySearch(intArray, 5, 0, len(intArray))
	if res != false {
		fmt.Printf("%d", res)
	} else {
		fmt.Printf("no find!")
	}
}

func binarySearch(array []int, i int, low int, high int) any {
	if low == high {
		return false
	} else {
		current := (low + high) / 2
		if array[current] == i {
			return current
		} else {
			if array[current] > i {
				return binarySearch(array, i, low, current-1)
			} else {
				return binarySearch(array, i, current+1, high)
			}
		}

	}
}
