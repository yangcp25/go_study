package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 利用cpu 多核加快计算速度
	test2()
}

func test2() {
	start := time.Now()

	// cpu 的数量 （物理核心指cpu有几个真正插在插槽上的数量，逻辑核心是指运行时可并行计算的cpu数量<指结合cpu多核和多线程技术得到的cpu个数>；这里指逻辑核心）
	cpuNums := runtime.NumCPU()

	fmt.Printf("我电脑的cpu逻辑核心数是%d\n", cpuNums)
	// 设置协程使用时可利用cpu数量
	runtime.GOMAXPROCS(cpuNums)

	chs := make([]chan int, cpuNums)

	for i := 0; i < cpuNums; i++ {
		chs[i] = make(chan int)
		go add2(i, chs[i])
	}

	total := 0
	for _, ch := range chs {
		sum := <-ch
		total += sum
	}
	end := time.Now()
	timeLong := end.Sub(start).Seconds()
	fmt.Println("花了这么久：\n", timeLong)
	fmt.Println("总共加起来是：\n", total)
}
func add2(seq int, ch chan int) {
	defer close(ch)
	sum := 0
	for i := seq; i < 1000000; i++ {
		sum += i
	}
	fmt.Printf("序号%d的计算结果是：%d\n", seq, sum)
	ch <- sum
}
