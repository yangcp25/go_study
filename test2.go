package main

import "fmt"

func main() {
	//

	a := []string{"1", "2", "2", "3"}

	a = removeSliceRepeat[string](a)

	fmt.Println(a)
}

func removeSliceRepeat[T comparable](array []T) []T {
	checkMap := make(map[T]struct{})
	tempArray := make([]T, 0)
	for _, v := range array {
		if _, ok := checkMap[v]; !ok {
			checkMap[v] = struct{}{}
			tempArray = append(tempArray, v)
		}
	}
	return tempArray
}
