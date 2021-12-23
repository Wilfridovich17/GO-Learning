/*
Race conditions
Wilfrido J. Paredes-Garcia
*/
package main

import (
	"fmt"
	"time"
)

var x int

//Goroutine 1
func update_x() {
	for t1 := 0; t1 <= 3; t1++ {
		x += 1
	}
}

//Goroutine 2
func print_x() {
	for t2 := 0; t2 <= x; t2++ {
		fmt.Println(x)
	}
}

/*
Main function calls two goroutines
I ran several times this code and
i got sometimes 0-4-4-4-4 and
and sometimes 4-4-4-4-4

This is a race condition because
we got different results depending
of interleavings.
*/
func main() {
	fmt.Println("!...Main Go-routine Start...!")
	// calling Goroutine 1
	go update_x()
	// calling Goroutine 2
	go print_x()
	//Time waiting to see goroutines
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("\n!...Main Go-routine End...!")
}
