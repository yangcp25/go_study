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
// 存在一个 m x n 的二维数组，其成员取值范围为0或1，其中值为1的成员具备扩散性，每经过1S，将上下左右值为0的成员同化为1。
// 二维数组的成员初始值都为0，将第[i,j]和[k,l]两个位置上元素修改成1后，求矩阵的所有元素变为1需要多长时间。
//
// 输入描述
//
// 输入数据中的前2个数字表示这是一个mxn的矩阵，m和n不会超过1024大小;
// 中间两个数字表示一个初始扩散点位置为 i,j
// 最后2个数字表示另一个扩散点位置为 k,l
//
// 输出描述
//
// 输出矩阵的所有元素变为1所需要秒数

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	strArr := strings.Split(str, ",")

	// 使用更具描述性的变量名
	rows, _ := strconv.Atoi(strArr[0])
	cols, _ := strconv.Atoi(strArr[1])
	r1, _ := strconv.Atoi(strArr[2])
	c1, _ := strconv.Atoi(strArr[3])
	r2, _ := strconv.Atoi(strArr[4])
	c2, _ := strconv.Atoi(strArr[5])

	// 特殊情况处理：如果矩阵只有一个或两个点，且初始点就是这些点，时间为0
	if rows*cols <= 2 {
		fmt.Println(0)
		return
	}

	// grid 用于记录每个格子的状态，0表示未被感染，1表示已被感染
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}

	// queue 用于实现广度优先搜索 (BFS)，存储待处理的坐标点
	// 使用一个 [2]int 来表示坐标，比转换成一维 key 更直观且不易出错
	queue := make([][2]int, 0)

	// 初始化两个扩散源
	if grid[r1][c1] == 0 {
		grid[r1][c1] = 1
		queue = append(queue, [2]int{r1, c1})
	}
	// 如果两个初始点是同一个，只添加一次
	if grid[r2][c2] == 0 {
		grid[r2][c2] = 1
		queue = append(queue, [2]int{r2, c2})
	}

	infectedCount := len(queue)
	totalCount := rows * cols
	time := 0

	// 定义四个方向的移动
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上, 下, 左, 右

	// 当还有未感染的格子，并且队列不为空时，继续循环
	for infectedCount < totalCount && len(queue) > 0 {
		time++

		// 当前层级的节点数量（在这一秒内需要处理的所有扩散源）
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			// 取出队列头部的节点
			curr := queue[0]
			queue = queue[1:]
			currR, currC := curr[0], curr[1]

			// 遍历四个方向
			for _, dir := range directions {
				nextR, nextC := currR+dir[0], currC+dir[1]

				// 检查新坐标是否越界，以及是否已经被感染过
				if nextR >= 0 && nextR < rows && nextC >= 0 && nextC < cols && grid[nextR][nextC] == 0 {
					grid[nextR][nextC] = 1 // 标记为已感染
					infectedCount++
					queue = append(queue, [2]int{nextR, nextC}) // 加入队列，作为下一秒的扩散源
				}
			}
		}
	}

	fmt.Println(time)
}
