package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	fmt.Printf("s1:%+v\n", s1)
	s2 := s1[2:6:7]
	fmt.Printf("s2:%+v\n", s2)
	s2 = append(s2, 100)
	fmt.Printf("s2 append:%+v\n", s2)

	s2 = append(s2, 200)
	fmt.Printf("s2 append: append%+v\n", s2)
	s1[2] = 20
	fmt.Printf("change s1[2] append: append%+v\n	", s2)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)

	testAppend()

	s := []int{1, 1, 1}
	testFunc(s)
	fmt.Println(s)
}

func testFunc(s []int) {
	// i只是一个副本，不能改变s中元素的值
	for _, i := range s {
		i++
	}
	for _, i := range s {
		print(i)
	}

	//for i := range s {
	//	s[i] += 1
	//}
}

func testAppend() {
	s := make([]int, 0)

	oldCap := cap(s)

	for i := 0; i < 2048; i++ {
		s = append(s, i)

		newCap := cap(s)

		if newCap != oldCap {
			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
			oldCap = newCap
		}
	}
}
