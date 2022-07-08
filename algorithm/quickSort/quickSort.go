package main

import "fmt"

// 选择
// 时间复杂度 O(n2)
// 稳定的算法
func main() {
	intArray := []int{4, 5, 6, 7, 3, 2, 1, 8, 0}
	intArray2 := []int{2, 1}
	quickSort(intArray[:], 0, len(intArray)-1)
	quickSort(intArray2[:], 0, len(intArray2)-1)
	fmt.Printf("%v", intArray)
	fmt.Printf("%v", intArray2)
}

func quickSort(arrayData []int, q int, p int) {
	if q >= p {
		return
	}

	r := sortHandle(arrayData[:], q, p)

	quickSort(arrayData[:], q, r-1)
	quickSort(arrayData[:], r+1, p)
}

func sortHandle(arrayData []int, q int, p int) int {
	i, j := q, q
	for j < p {
		// i不动 相当于表示比arrayData[p]小的位置
		if arrayData[j] < arrayData[p] {
			// 通过ij 交换将比arrayData[p]小的放到i处
			arrayData[i], arrayData[j] = arrayData[j], arrayData[i]
			i++
		}
		j++
	}
	// 最后比较的数 放在标识的位置
	arrayData[i], arrayData[p] = arrayData[p], arrayData[i]
	return i
}
