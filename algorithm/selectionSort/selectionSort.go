package main

import "fmt"

// 选择
// 时间复杂度 O(n2)
// 不稳定的算法
func main() {
	intArray := []int{4, 5, 6, 7, 8, 3, 2, 1}
	selectionSort(intArray[:])
	fmt.Printf("%v", intArray)
}

func selectionSort(arrayData []int) {
	length := len(arrayData)
	for i := 0; i < length; i++ {
		min := i
		j := i + 1
		for ; j < length; j++ {
			if arrayData[min] > arrayData[j] {
				min = j
			}
		}
		if min != j {
			arrayData[i], arrayData[min] = arrayData[min], arrayData[i]
		}
	}
}
