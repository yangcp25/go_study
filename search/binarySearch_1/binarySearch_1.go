package main

import (
	"fmt"
	"sort"
)

// 查找算法
// 二分查找（）
func main() {
	intArray := []int{4, 6, 5, 3, 1, 2, 8, 2, 7}
	sort.Ints(intArray)
	fmt.Printf("%d", intArray)
	res := binarysearch1(intArray, 2, 0, len(intArray))
	if res != false {
		fmt.Printf("%d", res)
	} else {
		fmt.Printf("no find!")
	}
}

func binarysearch1(array []int, i int, low int, high int) any {
	if low == high {
		return false
	} else {
		current := (low + high) / 2
		if array[current] > i {
			return binarysearch1(array, i, low, current-1)
		} else if array[current] < i {
			return binarysearch1(array, i, current+1, high)
		} else {
			if array[current] == i || array[current-1] != i {
				return current
			} else {
				return binarysearch1(array, i, low, current-1)
			}
		}

	}
}
