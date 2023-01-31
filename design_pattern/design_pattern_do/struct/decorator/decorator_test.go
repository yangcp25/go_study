package decorator

import (
	"testing"
)

func TestNewColorSquare(t *testing.T) {
	square := square{}
	drawHandler := newColorSquare(square, "co")

	drawHandler.Draw()
}
