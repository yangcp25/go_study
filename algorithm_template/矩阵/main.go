package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	res := make([]int, 0)

	for top <= bottom && left <= right {
		// 1. 从左到右
		for col := left; col <= right; col++ {
			res = append(res, matrix[top][col])
		}
		top++

		// 2. 从上到下
		for row := top; row <= bottom; row++ {
			res = append(res, matrix[row][right])
		}
		right--

		// 防止越界再走
		if top > bottom || left > right {
			break
		}

		// 3. 从右到左
		for col := right; col >= left; col-- {
			res = append(res, matrix[bottom][col])
		}
		bottom--

		// 4. 从下到上
		for row := bottom; row >= top; row-- {
			res = append(res, matrix[row][left])
		}
		left++
	}

	return res
}
