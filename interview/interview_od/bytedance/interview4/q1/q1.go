package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 实现一个函数 calTotalPrice，该函数根据 实际用量（realQuantity）、用量阶梯（quantityInterval） 和 单价阶梯（unitPriceInterval） 来计算总价。
// 用量阶梯和单价阶梯是用逗号分隔的字符串，表示不同的区间和对应的单价。
// 输入输出示例
// ● 用量阶梯（quantityInterval）: "0,100,200,300"
// ● 单价阶梯（unitPriceInterval）: "50,30,20,10"
// 含义：
// ● [0, 100) 个，单价 50 元
// ● [100, 200) 个，单价 30 元
// ● [200, 300) 个，单价 20 元
// ● [300, +∞) 个，单价 10 元
// 示例计算：
// 1. realQuantity = 51，总价 = 51(数量) * 50(单价) = 2550
// 2. realQuantity = 200，总价 = 200 * 20 = 4000
// 3. realQuantity = 500，总价 = 500 * 10 = 5000
func main() {
	realQuantity := ToInt("51,200,500")
	quantityInterval := ToInt("0,100,200,300")
	unitPriceInterval := ToInt("50,30,20,10")

	for _, quantity := range realQuantity {
		sum := calTotalPrice(quantity, quantityInterval, unitPriceInterval)
		fmt.Println(sum)
	}
}

func ToInt(realQuantityStr string) []int {
	realQuantityArray := strings.Split(realQuantityStr, ",")
	realQuantity := make([]int, 0)
	for _, v := range realQuantityArray {
		temp, _ := strconv.Atoi(v)
		realQuantity = append(realQuantity, temp)
	}

	return realQuantity
}

func calTotalPrice(total int, quantityInterval []int, unitPriceInterval []int) (price int) {

	key := 0
	for i := len(quantityInterval) - 1; i >= 0; i-- {
		if total >= quantityInterval[i] {
			key = i
			break
		}
	}

	price = total * unitPriceInterval[key]

	return price
}
