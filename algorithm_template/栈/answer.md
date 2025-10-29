
``` go
// 20 有效的括号
func isValid(s string) bool {
	pairs := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	stack := make([]rune, 0)
	for _, c := range s {
		switch c {
		case '}', ')', ']':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
// 739 每日温度
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
```