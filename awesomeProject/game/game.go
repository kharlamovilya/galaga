package game

import (
	"fmt"
)

type Vector2D struct {
	X, Y float64
}

func (v Vector2D) Plus(v1 Vector2D) Vector2D {
	return Vector2D{v.X + v1.X, v.Y + v1.Y}
}

func (v Vector2D) String() {
	fmt.Println("x: ", v.X, ", y: ", v.Y)
}
func Start() {
}
