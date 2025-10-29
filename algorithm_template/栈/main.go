package main

func main() {

}

// 20 有效的括号
func isValid(s string) bool {

	return true
}

// 739 每日温度
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[preIndex] = i - preIndex
		}
		stack = append(stack, i)
	}
	return res
}
