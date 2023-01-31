package decorator

import "fmt"

type IDraw interface {
	Draw()
}

type square struct{}

func (receiver square) Draw() {
	// todo
	fmt.Println("画了个正方形")
}

type colorSquare struct {
	square square
	color  string
}

func newColorSquare(square2 square, color string) *colorSquare {
	return &colorSquare{
		square: square2,
		color:  color,
	}
}

func (c colorSquare) Draw() {
	c.square.Draw()
	fmt.Println("并且是", c.color)
}
