package main

import "fmt"

func main() {
	s3 := make([]int, 2, 10)
	fmt.Println(s3) // [0 0]
	Test2(s3)
	fmt.Println(s3) // [0 0]
	fmt.Printf("s3 addr: %p\n", s3)

	s4 := s3[0:10]
	s4[0] = 100
	fmt.Println(s3)                  // [100 0]
	fmt.Printf("s3 addr: %+v\n", s3) // [100 0]
	fmt.Println(s4)                  // [100 0 6 6 6 0 0 0 0 0]
	fmt.Printf("s3 addr: %p, s4 addr: %p\n", s3, s4)
}
func Test2(s []int) {
	test := make([]int, 2)
	test = append(test, 1)
	test = append(test, 1)
	fmt.Printf("Test2 s addr: %p\n", s)
	s = append(s, 6)
	s = append(s, 6)
	s = append(s, 6)
	fmt.Println(s) // [0 0 6 6 6]
	fmt.Printf("Test2 append s addr: %p\n", s)
}
