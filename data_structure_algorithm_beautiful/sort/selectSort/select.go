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

	selectSort(testArray)

	fmt.Printf("排序后的数组为：%+v\n", testArray)

	intArray := []int{4, 5, 6, 7, 8, 3, 2, 1}

	selectSort(intArray)

	fmt.Printf("排序后的数组为：%+v", intArray)
}

func selectSort[T constraints.Ordered](data []T) {
	length := len(data)
	for i := 0; i < length; i++ {
		minNum := data[i]
		minKey := i
		for j := i + 1; j < length; j++ {
			if minNum > data[j] {
				minNum = data[j]
				minKey = j
			}
		}
		data[i], data[minKey] = data[minKey], data[i]
	}
}
