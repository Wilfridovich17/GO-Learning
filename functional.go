package main

import "fmt"

func increase(number int) int {
	return number + 1
}

func applyFUNC(a_func func(int) int, value int) int {
	return a_func(value)
}

func main() {
	var a_number int = 5
	var the_func func(int) int

	the_func = increase
	fmt.Printf("%d", the_func(a_number))

	fmt.Printf("%d", applyFUNC(increase, 5))
	fmt.Printf("%d", applyFUNC(func(x int) int { return x + 1 }, 5))
}
