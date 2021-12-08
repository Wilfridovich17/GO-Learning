package main

import (
	"fmt"
	"math"
)

func origindistance(o_x float64, o_y float64) func(float64, float64) float64 {
	new_function := func(x float64, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(x-o_y, 2))
	}

	return new_function
}

func main() {
	dist_classic_origin := origindistance(0, 0)
	dist_new_origin := origindistance(2, 2)

	fmt.Printf("%f\n", dist_classic_origin(2, 2))
	fmt.Printf("%f\n", dist_new_origin(2, 2))
}
