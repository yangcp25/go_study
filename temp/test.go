package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var str string
	//Scan(&str)
	//fmt.Println(getLastWordLenght(str))
	test2()
}

func getLastWordLenght(str string) (ans int) {
	strSlice := strings.Split(str, " ")
	return len(strSlice[len(strSlice)-1])
}

func Scan(str *string) {
	input := bufio.NewReader(os.Stdin)
	data, _, _ := input.ReadLine()
	*str = string(data)
}

func test2() {
	getInputs()

}

func getInputs() {
	a := 0
	b := 0
	fmt.Scan(&a, &b)

	fmt.Println(a, b)
}
