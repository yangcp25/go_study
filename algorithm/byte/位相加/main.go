package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := "11111"
	b := "10111"
	res := addBinary(a, b)
	fmt.Println(res)
}

func addBinary(a string, b string) string {
	x := new(big.Int)
	y := new(big.Int)

	x.SetString(a, 2)
	y.SetString(b, 2)

	res := new(big.Int).Add(x, y)

	return res.Text(2)
}

func addBinaryV2(a string, b string) string {

}
