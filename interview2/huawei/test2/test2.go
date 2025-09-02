package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 题目描述
//
// 存在一个mxn的二维数组，其成员取值范围为0或1，其中值为1的成员具备扩散性，每经过1S，将上下左右值为0的成员同化为1，二维数组的成员初始值都为0，将第[i,j]和[k,l]两个个位置上元素修改成1后，求矩阵的所有，元素变为1需要多长时间
//
// 输入描述
//
// 输入数据中的前2个数字表示这是一个mxn的矩阵，m和n不会超过1024大小;
//
// 中间两个数字表示一个初始扩散点位置为I,j
//
// 最后2个数字表示另一个扩散点位置为k,l
//
// 输出描述
//
// 输出矩阵的所有元素变为1所需要秒数
//
// 用例
func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	strArr := strings.Split(str, ",")
	m, n := Atoi(strArr[0]), Atoi(strArr[1])
	pos1, pos2, pos3, pos4 := Atoi(strArr[2]), Atoi(strArr[3]), Atoi(strArr[4]), Atoi(strArr[5])

	dp := make([][]int, m)
	for k := range dp {
		dp[k] = make([]int, n)
	}
	dp[pos1][pos2] = 1
	dp[pos3][pos4] = 1
	//fmt.Println(dp[pos1][pos2])
	//fmt.Println(dp[pos3][pos4])
	//count := m*n
	path := make([]int, 0)
	path = append(path, GetPosKey(pos1, pos2, m), GetPosKey(pos3, pos4, m))
	//fmt.Println(path)
	times, count := 0, 2
OutLoop:
	for len(path) > 0 {
		times++
		length := len(path)
		fmt.Println("path", path)
		temp := make([]int, 0)
		//copy(temp, path)
		for i := 0; i < length; i++ {
			// 取得上下左右
			beihor := GetNeibor(path[i], m, n)
			fmt.Println(path, path[i], "beihor", beihor)
			for _, v := range beihor {
				posT1, posT2 := GetPos(v, n)
				if dp[posT1][posT2] == 0 {
					temp = append(temp, v)
					count++
					fmt.Println("count", count)
					dp[posT1][posT2] = 1
					fmt.Println(dp)
					if count == m*n {
						break OutLoop
					}
				}
			}
		}
		path = temp
	}
	fmt.Println(times)

}

func Atoi(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}

func GetPosKey(i, j, m int) int {
	return m*i + j
}

// 5 x 4
func GetPos(i, n int) (int, int) {
	row := i / n
	col := i - row*n
	return row, col
}

func GetNeibor(i, m, n int) []int {
	t1, t2 := GetPos(i, n)
	neibor := make([]int, 0)
	// 上
	if t2-1 >= 0 {
		//neibor = append(neibor, t1-1)
		neibor = append(neibor, GetPosKey(t1, t2-1, m))
	}
	if t2+1 <= m-1 {
		//neibor = append(neibor, t2+1)
		neibor = append(neibor, GetPosKey(t1, t2+1, m))
	}

	if t1-1 >= 0 {
		//neibor = append(neibor, t1-1)
		neibor = append(neibor, GetPosKey(t1-1, t2, m))
	}

	if t1+1 <= n-1 {
		neibor = append(neibor, GetPosKey(t1+1, t2, m))
	}
	return neibor
}
