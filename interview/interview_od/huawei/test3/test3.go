package test3

import (
	"fmt"
	"unicode"
)

func main() {

	//输入一个字符串，里面有大写字母、小写字母、数字，
	//处理结束后，大写字母在前，然后是小写字母，最后是数字，
	//要求：在原有字符串上做交换实现，不要建新的数据结构。
	//mN1oO2pP3
	//Nm1oO2pP3
	str := []rune("mN1oO2pP3")
	// AbcafafaaFAs
	exchangePos(str)

	fmt.Println(string(str))
}

func exchangePos(str []rune) {
	// i 找非大写字母 j 找大写字母 如果 i ！= j 交换
	length := len(str)
	i, j := 0, 0
	for i < length && j < length {
		if isUper(str[j]) {
			//&& !isUper(rune(str[i]))
			if i != j {
				str[i], str[j] = str[j], str[i]
				i++
			}
		}
		if isUper(str[i]) {
			i++
		}
		j++
	}
	// i 从上一次结束为止开始，找 数字， j 找小写字母 如果 i ！= j 交换
	j = i
	//fmt.Println(string(str), i, j)
	for i < length && j < length {
		if isLower(str[j]) {
			//&& !isUper(rune(str[i]))
			if i != j {
				str[i], str[j] = str[j], str[i]
				i++
			}
		}
		if isLower(str[i]) {
			i++
		}
		j++
	}
	return
}

func isUper(s rune) bool {
	return unicode.IsUpper(s)
}
func isLower(s rune) bool {
	return unicode.IsLower(s)
}
