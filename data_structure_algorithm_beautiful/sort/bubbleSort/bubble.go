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

	bubbleSort(testArray)

	fmt.Printf("排序后的数组为：%+v", testArray)
}

func bubbleSort[T constraints.Ordered](data []T) {
	length := len(data)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}
