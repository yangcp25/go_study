package main

import "fmt"

// 选择
// 时间复杂度 O(n2)
// 稳定的算法
func main() {
	intArray := []int{4, 5, 6, 7, 3, 2, 1, 8, 9}
	intArray2 := []int{4, 5, 6, 7, 8, 3, 2, 1}
	intArray3 := []int{2, 1}
	quickSort(intArray[:], 0, len(intArray)-1)
	quickSort(intArray2, 0, len(intArray2)-1)
	quickSort(intArray3, 0, len(intArray3)-1)
	fmt.Printf("%v\n", intArray)
	fmt.Printf("%v\n", intArray2)
	fmt.Printf("%v", intArray3)
}

func quickSort(arrayData []int, start, end int) {
	if start >= end {
		return
	}

	pivot := quickSortDo(arrayData, start, end)

	quickSort(arrayData, start, pivot-1)
	quickSort(arrayData, pivot+1, end)
}

// 将当前区间 以end 为标准 分成 前后2个区间 前面的区间小于 end ，后面的end 大于end
func quickSortDo(data []int, start, end int) int {
	// 定义初始位置
	startSmall, startLarge := start, start
	for startLarge < end {
		// startSmall, startLarge 2个标识 分别标识 小于 end 的 和大于end 的位置
		if data[startLarge] < data[end] {
			if startSmall != startLarge {
				data[startSmall], data[startLarge] = data[startLarge], data[startSmall]
			}
			startSmall++
		}
		startLarge++
	}
	data[startSmall], data[end] = data[end], data[startSmall]
	// startSmall 相当于标识一个中间数
	return startSmall
}
