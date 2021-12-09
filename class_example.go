package main

import "fmt"

type math_func func(float64) float64

func (f math_func) Tabulate(x_start, x_end float64, n int) ([]float64, []float64) {
	step := (x_end - x_start) / float64(n)
	aux := x_start
	x := make([]float64, 0)
	y := make([]float64, 0)
	for aux < x_end {
		x = append(x, aux)
		y = append(y, f(aux))
		aux += step
	}
	return x, y
}

func main() {
	squared := math_func(func(a float64) float64 { return a * a })
	_, t_squared := squared.Tabulate(0, 10, 10)
	for idx, f_t := range t_squared {
		fmt.Printf("%d --- %f\n", idx, f_t)
	}
}
