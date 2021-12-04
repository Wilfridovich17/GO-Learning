package main

import "fmt"

func main() {
	type Persona struct {
		name string
		age  int
	}

	my_struct := Persona{name: "William", age: 56}

	fmt.Println(my_struct.name, my_struct.age)
}
