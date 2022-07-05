package main

import "fmt"

// 冒泡排序 （思想就是通过2个相邻元素的交换 将最大或者最小放到最末尾或者最开始)
// 时间复杂度 O(n2)
// 相邻元素不会发生交换 稳定的算法
func main() {
	intArray := [8]int{4, 5, 6, 7, 8, 3, 2, 1}
	res := bubbleSort(intArray[:])
	fmt.Printf("%v", res)
}

func bubbleSort(arrayData []int) []int {
	length := len(arrayData)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arrayData[i] > arrayData[j] {
				arrayData[i], arrayData[j] = arrayData[j], arrayData[i]
			}
		}
	}
	return arrayData
}
