package main

import (
	"GO-Learning/vector"
	"fmt"
)

func main() {
	var v vector.Vec2d
	v.InitV(2, 3)
	fmt.Printf("%f", v.Norm())
}
