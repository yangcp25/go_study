
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
```go
import "strconv"
// 逆波兰表达式（Evaluate Reverse Polish Notation）
func evalRPN(tokens []string) int {
    stack := []int{}
    for _, t := range tokens {
        switch t {
        case "+", "-", "*", "/":
            b := stack[len(stack)-1]
            a := stack[len(stack)-2]
            stack = stack[:len(stack)-2]
            var res int
            switch t {
            case "+": res = a + b
            case "-": res = a - b
            case "*": res = a * b
            case "/": res = a / b
            }
            stack = append(stack, res)
        default:
            num, _ := strconv.Atoi(t)
            stack = append(stack, num)
        }
    }
    return stack[0]
}

```