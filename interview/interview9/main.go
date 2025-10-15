package main

import "fmt"

func main() {
	// <问题描述>
	//给定一个包含n个整数元素的数组a[i]，满足a[i] <= a[j]，0 <= i < j < n
	//请你统计大小为k的元素的个数并返回，如果不存在则返回0。
	//效率越高的算法将有助于获得更高的分数。
	//<输入>
	//第1行包含两个整数n和k，表示数组元素个数和要统计的元素大小为k
	//第2行包含n个数，表示数组的每个元素
	//<输出>
	//一个整数，即大小为k的元素的个数
	//输入样例	输出样例
	//7 5
	//3 4 5 5 5 7 9
	/**
	7 5
	3 4 5 5 5 7 9
	3
	*/

	nums := []int{3, 4, 5, 5, 5, 7, 9}
	nums2 := 5
	count := 0
	for _, num := range nums {
		if num == nums2 {
			count++
		} else if num > nums2 {
			break
		}
	}
	fmt.Println(count)
}
