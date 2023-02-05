package generics

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	v, ok := intStack.Pop()
	fmt.Println(v, ok)
	//intStack.Push("nope")
}
