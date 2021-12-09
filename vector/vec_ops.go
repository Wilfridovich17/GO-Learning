package vector

import "math"

type Vec2d struct {
	coord_x float64
	coord_y float64
}

func (vec *Vec2d) InitV(x, y float64) {
	vec.coord_x = x
	vec.coord_y = y
}

func (vec *Vec2d) Norm() float64 {
	return math.Sqrt(math.Pow(vec.coord_x, 2) + math.Pow(vec.coord_y, 2))
}
