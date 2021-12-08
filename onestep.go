package main

import "fmt"

func diffeqn_example(x float64, t float64) float64 {
	return -0.5*x + 0*t
}

func one_step_method(f func(u float64, t float64) float64, h float64, method string) func(float64, float64) float64 {
	var gen_func func(float64, float64) float64

	switch method {
	case "euler":
		gen_func = func(u_i float64, t_i float64) float64 {
			return u_i + h*f(u_i, t_i)
		}
	case "midpoint":
		gen_func = func(u_i float64, t_i float64) float64 {
			return u_i + h*f(u_i+(h/2.0)*f(u_i, t_i), t_i+h/2.0)
		}
	case "rk4":
		gen_func = func(u_i float64, t_i float64) float64 {
			s1 := f(u_i, t_i)
			s2 := f(u_i+(h/2)*s1, t_i+h/2)
			s3 := f(u_i+(h/2)*s2, t_i+h/2)
			s4 := f(u_i+s3, t_i+h)

			return u_i + h*(s1+2*s2+2*s3+s4)/6
		}
	}
	return gen_func
}

func main() {
	var n_points int = 20
	var aux_t float64 = 0.0

	aux_x := [3]float64{15.0, 15.0, 15.0}

	x_te := []float64{15.0}
	x_tm := []float64{15.0}
	x_tr := []float64{15.0}
	t := []float64{0.0}

	euler := one_step_method(diffeqn_example, 0.5, "euler")
	midpoint := one_step_method(diffeqn_example, 0.5, "midpoint")
	rk := one_step_method(diffeqn_example, 0.5, "rk4")

	for i := 0; i < n_points; i++ {
		aux_x[0] = euler(aux_x[0], aux_t)
		x_te = append(x_te, aux_x[0])

		aux_x[1] = midpoint(aux_x[1], aux_t)
		x_tm = append(x_tm, aux_x[1])

		aux_x[2] = rk(aux_x[2], aux_t)
		x_tr = append(x_tr, aux_x[2])

		t = append(t, aux_t+0.5)
		aux_t = aux_t + 0.5
	}

	fmt.Println("Solutions Euler:")
	for idx, value := range x_te {
		fmt.Printf("%d -- %f\n", idx, value)
	}

	fmt.Println("\nSolutions Midpoint:")
	for idx, value := range x_tm {
		fmt.Printf("%d -- %f\n", idx, value)
	}

	fmt.Println("\nSolutions RK4:")
	for idx, value := range x_tr {
		fmt.Printf("%d -- %f\n", idx, value)
	}
}
