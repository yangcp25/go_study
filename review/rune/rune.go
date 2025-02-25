package main

import "fmt"

func main() {
	//test := []byte("中国")
	//test2 := []rune("中国")
	//fmt.Println(test)
	//fmt.Println(test2)
	//test3 := []byte{1, 2, 3}
	//fmt.Println(test3)
	//
	//f, err := os.Create("test.text")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//_, err = f.Write([]byte{65, 66, 67})
	//if err != nil {
	//	fmt.Println(err)
	//}

	//s := "你mmp"
	//
	//for i := 0; i < len(s); i++ {
	//	fmt.Println(string(s[i]))
	//}
	//
	//fmt.Println("---")
	//for _, v := range s {
	//	fmt.Println(v)
	//}

	// 错误的写法
	// 错误的写法
	fmt.Println("错误的写法:")
	funcs1 := make([]func(), 3)
	for i := 0; i < 3; i++ {
		funcs1[i] = func() {
			fmt.Println(i)
		}
	}
	// 执行函数
	for _, f := range funcs1 {
		f()
	}

	fmt.Println("\n正确的写法:")
	funcs2 := make([]func(), 3)
	for i := 0; i < 3; i++ {
		i := i // 创建新的变量
		funcs2[i] = func() {
			fmt.Println(i)
		}
	}
	// 执行函数
	for _, f := range funcs2 {
		f()
	}
}
