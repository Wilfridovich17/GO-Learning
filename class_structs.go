package main

import "fmt"

type math_func struct {
	f   func(float64) float64
	l_i float64
	l_u float64
}

func (F math_func) Tabulate(n int) ([]float64, []float64) {
	step := (F.l_u - F.l_i) / float64(n)
	aux := F.l_i
	x := make([]float64, 0)
	y := make([]float64, 0)
	for aux < F.l_u {
		x = append(x, aux)
		y = append(y, F.f(aux))
		aux += step
	}
	return x, y
}

func main() {
	squared := math_func{f: func(a float64) float64 { return a * a }, l_i: 0, l_u: 10}
	_, t_squared := squared.Tabulate(10)
	for idx, f_t := range t_squared {
		fmt.Printf("%d --- %f\n", idx, f_t)
	}
}
