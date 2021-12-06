package main

import "fmt"

//This function has a call by value
func example(entero int) int {
	return entero + 1
}

//This function has a call by reference
func pointer_example(entero *int) *int {
	*entero = *entero + 1
	return entero
}

func main() {
	var x int = 2
	var y int = 2
	//Results of calling by value
	fmt.Println("Call by value:\n")
	fmt.Printf("original: %d\nfunction result: %d\nafter function: %d", x, example(x), x)
	//Results of calling by reference
	fmt.Println("\n\nCall by reference:\n")
	fmt.Printf("original: %d\n", y)
	fmt.Printf("result: %d\nafter function: %d", *pointer_example(&y), y)

}
