package main

import "fmt"

// 冒泡排序 （思想就是通过2个相邻元素的交换 将最大或者最小放到最末尾或者最开始)
// 时间复杂度 O(n2)
// 相邻元素不会发生交换 稳定的算法
func main() {
	intArray := []int{4, 5, 3}
	res := bubbleSort(intArray[:])
	fmt.Printf("%v", res)
}

func bubbleSort(arrayData []int) []int {
	length := len(arrayData)
	for i := 0; i < length; i++ {
		j := i - 1
		val := arrayData[i]
		for ; j >= 0; j-- {
			if val > arrayData[j] {
				break
			} else {
				arrayData[j+1] = arrayData[j]
			}
		}
		arrayData[j+1] = val
	}
	return arrayData
}
