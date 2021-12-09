package main

import (
	"fmt"
	"math"
)

func assign_values(msg string, value *float64) {
	for {
		fmt.Printf(msg)
		_, err := fmt.Scan(value)
		if err == nil {
			break
		}
	}
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	f := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}

	return f
}

func main() {
	var user_input [4]float64
	msgs := [4]string{"Acceleration: ", "Initial velocity: ", "Initial displacement:", "Time: "}

	for i := 0; i < 4; i++ {
		assign_values(msgs[i], &user_input[i])
	}

	user_displaceFN := GenDisplaceFn(user_input[0], user_input[1], user_input[2])

	fmt.Printf("The displacement at time %f is %f:", user_input[3], user_displaceFN(user_input[3]))
}
