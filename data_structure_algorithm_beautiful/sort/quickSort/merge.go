package main

import "fmt"

// 选择
// 时间复杂度 O(n2)
// 稳定的算法
func main() {
	intArray := []int{4, 5, 6, 7, 8, 3, 2, 1}
	intArray = quickSort(intArray[:])
	fmt.Printf("%v", intArray)
}

func quickSort(arrayData []int) []int {
	length := len(arrayData)
	if length <= 1 {
		return arrayData
	}

	count := length / 2

	left := quickSort(arrayData[:count])
	right := quickSort(arrayData[count:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	i, j := 0, 0
	leftLength := len(left)
	rightLength := len(right)
	var res []int

	//
	for {
		if i >= leftLength || j >= rightLength {
			break
		}
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	if i < leftLength {
		for ; i < leftLength; i++ {
			res = append(res, left[i])
		}
	}

	if j < rightLength {
		for ; j < rightLength; j++ {
			res = append(res, right[j])
		}
	}
	return res
}
