package main

import "fmt"
import "golang.org/x/exp/constraints"

func main() {
	testArray := make([]int, 0, 10)
	testArray = append(testArray, 1)
	testArray = append(testArray, 6)
	testArray = append(testArray, 5)
	testArray = append(testArray, 7)
	testArray = append(testArray, 4)
	testArray = append(testArray, 4)
	testArray = append(testArray, 8)
	testArray = append(testArray, 3)
	testArray = append(testArray, 2)
	testArray = append(testArray, 9)

	insertSort(testArray)

	fmt.Printf("排序后的数组为：%+v\n", testArray)

	intArray := []int{4, 5, 6, 7, 8, 3, 2, 1}

	insertSort(intArray)

	fmt.Printf("排序后的数组为：%+v", intArray)
}

func insertSort[T constraints.Ordered](data []T) {
	length := len(data)
	for i := 1; i < length; i++ {
		if data[i-1] > data[i] {
			temp := data[i]
			// 找到插入的位置 比如 1342 42这里 只需要从4往前找到它应该插入的位置
			j := i
			for ; j > 0; j-- {
				if data[j-1] > temp {
					data[j] = data[j-1]
				} else {
					break
				}

			}
			data[j] = temp
		}

	}
}
